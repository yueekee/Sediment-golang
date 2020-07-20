package main

import (
	"fmt"
	"strings"
)

func main() {
	info := "https://chhcollege-static.oss-cn-shanghai.aliyuncs.com/xcxmatrix/gongkaike/0902霍夫曼/0902霍夫曼2.史蒂夫 霍夫曼.mp4"
	//info2 := "https://chhcollege-static.oss-cn-shanghai.aliyuncs.com/xcxmatrix/gongkaike/霍夫曼预告1213.mp4"

	left := strings.TrimLeft(info, "https://chhcollege-static.oss-cn-shanghai.aliyuncs.com/xcxmatrix/gongkaike/")
	splitUrl := strings.Split(left, "/")
	if len(splitUrl) >=2 {

	}
	fmt.Println("+++++len", len(splitUrl))
	fmt.Println("+++++splitUrl", splitUrl)


	info = "https://static.chuanghehui.com/xcxmatrix/gongkaike/" + splitUrl[len(splitUrl)-1]
	fmt.Println("+++++info", info)
}
