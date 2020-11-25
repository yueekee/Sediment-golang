package main

import (
	"fmt"
	"strings"
)

// TrimLeft和Trim分割字符串时，对于有"/"的字符串切割会出现问题
// 切割前缀使用TrimPrefix，切割后缀使用TrimSuffix进行代替
func main() {
	Url := "https://www.baidu.com/material/list"
	fmt.Println(strings.TrimLeft(Url, "https://www.baidu.com"))   // erial/list
	fmt.Println(strings.TrimPrefix(Url, "https://www.baidu.com")) // /material/list

	Url2 := ":www.baidu.com/material/list"
	fmt.Println(strings.Trim(Url, "https://www.baidu.com")) // erial/l
	fmt.Println(strings.Trim(Url2, ":www.baidu.com"))       // /material/list

	s := "Goodbye,, world!"
	fmt.Println(strings.TrimSuffix(s, ", world!")) // Goodbye,
}
