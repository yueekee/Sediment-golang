package main

import "fmt"

/* 使用命令：
go tool compile -S 11-slicecompile.go
*/

func main() {
	slice := make([]int, 5, 10) // 长度为5，容量为10
	slice[2] = 2 // 索引为2的元素赋值为2
	fmt.Println(slice)
}
