package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "a156c a56c a1c ac auc"
	reg := regexp.MustCompile(`a[1u]c`) //返回正则表达式
	result := reg.FindAllString(str, -1) //查找所有符合规则的字符串，-1表示匹配全部
	fmt.Println(result)	// [a1c auc]
}
