package main

import (
	"fmt"
	"os"
)

// os.Args是一个字符串切片，代表了当前执行程序的命令行参数。
// go build 02-Args获取命令行参数.go
// ./02-Args获取命令行参数 123 1.1
func main() {
	fmt.Println(os.Args)     // [./02-Args获取命令行参数 123 1.1]
	fmt.Println(os.Args[1:]) // [123 1.1]
}
