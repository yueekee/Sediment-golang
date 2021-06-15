package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(`http://172.16.1.120:8080/down?pid=123412341&key=1234123412`, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("sample.pdf", buf, 0644); err != nil {
		log.Fatal(err)
	}
}

// print a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			pdfParams := page.PrintToPDF()
			pdfParams.Landscape = false              // 横向打印
			pdfParams.PrintBackground = true         // 打印背景图.  默认false.
			pdfParams.PreferCSSPageSize = true       // 是否首选css定义的页面大小？默认false,将自动适应.
			pdfParams.IgnoreInvalidPageRanges = true // 是否要忽略非法的页码范围. 默认false.
			pdfParams.PaperWidth = 20.92             // 页面宽度(英寸). 默认8.5英寸.（24英寸 20.92 x 11.77）
			pdfParams.PaperHeight = 11.77            // 页面高度(英寸). 默认11英寸
			fmt.Printf("----params:%#v\n", pdfParams)
			buf, _, err := pdfParams.Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}