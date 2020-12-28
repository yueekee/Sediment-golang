package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Title("hello world")) // Hello World 将字符串所包含的每个单词的第一个字母大写

	var width int = 4
	var length float64 = 7.1

	fmt.Println("area is", width*int(length))
	fmt.Println("area2 is", float64(width)*length)
}
