package main

import (
	"fmt"
	"strings"
)

// TrimSpace 删除字符串开头和结尾的所有空白字符（换行符、制表符和常规空格）
func main() {
	info := "\t \n 史蒂夫 霍夫曼.mp4 \n"
	fmt.Println(strings.TrimSpace(info))	// 史蒂夫 霍夫曼.mp4
}
