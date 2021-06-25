package demo

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"

	"github.com/gin-gonic/gin"
)

// 下载报告

type ReportStr struct {
	Report string
}

type DownReportParam struct {
	PID      string `json:"pid"`
	Key      string `json:"key"`
	FileName string `json:"fileName"`
	Type     int    `json:"type"`
}

func GetPDF(c *gin.Context) {
	var (
		timenow         = fmt.Sprintf("%d", time.Now().Unix())
		prefixPath      = "report/" + timenow
		downReportParam DownReportParam
	)

	if err := c.BindJSON(&downReportParam); err != nil {
		fmt.Println("参数错误")
		return
	}

	pID := downReportParam.PID
	key := downReportParam.Key
	fileName := downReportParam.FileName
	// pdf路径，文件名
	pdfPath := prefixPath + ".pdf"
	err := os.MkdirAll("report/", 0777)
	if err != nil {
		c.String(http.StatusBadRequest, "文件夹创建失败")
		return
	}
	var url string
	if downReportParam.Type == 0 {
		url = fmt.Sprintf("%sdown?pid=%s&key=%s", "http://localhost:8080/", pID, key)
	}
	println(url)

	if err := ChromedpPrintPdf(url, pdfPath); err != nil {
		log.Println(err)
		fmt.Println(c, "下载失败")
		return
	}

	// pdf文件上传给前端
	bytes, err := ioutil.ReadFile(pdfPath)
	//defer func() {
	//	bytes = nil
	//}()

	//f, err := os.Open(pdfPath)
	//defer f.Close()
	//
	//if err != nil {
	//	log.Println("err1:", err)
	//}
	//
	//reader := bufio.NewReader(f)
	//var chunk []byte
	//buf := make([]byte, 1024)
	//for {
	//	buf, err = reader.ReadBytes('\n') //读取到\n结束
	//	if err != nil {
	//		if err == io.EOF { //文件已经结束
	//			break
	//		}
	//	}
	//	chunk = append(chunk, buf...)
	//}
	//defer func() {
	//	chunk = chunk[:0]
	//}()
	////fmt.Printf("chunk:%#v\n", string(chunk))
	//fmt.Println("------------len:", len(chunk))

	c.Header("content-disposition", `attachment; filename=`+fileName)
	c.Data(200, "application/pdf", bytes)
	c.String(http.StatusOK, "下载文件成功")
}

func ChromedpPrintPdf(url string, to string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	buf := make([]byte, 2048)
	err := chromedp.Run(ctx, printToPDF(url, &buf))
	if err != nil {
		return fmt.Errorf("chromedp Run failed,err:%+v", err)
	}

	if err := ioutil.WriteFile(to, buf, 0644); err != nil {
		return fmt.Errorf("write to file failed,err:%+v", err)
	}
	//buf = buf[:0] 不行
	return nil
}

// print a specific pdf page.
func printToPDF(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.WaitVisible(`.title-nav`, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfParams := page.PrintToPDF()
			pdfParams.Landscape = false              // 横向打印
			pdfParams.PrintBackground = true         // 打印背景图.  默认false.
			pdfParams.PreferCSSPageSize = true       // 是否首选css定义的页面大小？默认false,将自动适应.
			pdfParams.IgnoreInvalidPageRanges = true // 是否要忽略非法的页码范围. 默认false.
			pdfParams.PaperWidth = 20.92             // 页面宽度(英寸). 默认8.5英寸.（24英寸 20.92 x 11.77）
			pdfParams.PaperHeight = 11.77            // 页面高度(英寸). 默认11英寸
			buf, _, err := pdfParams.Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
