package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// os.Args是一个字符串切片，代表了当前执行程序的命令行参数。
// go build 02-Args获取命令行参数.go
// ./02-Args获取命令行参数 123 1.1
func main() {
	fmt.Println(os.Args)     // [./02-Args获取命令行参数 123 1.1]

	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}

	join1(os.Args)
	join2(os.Args)	// 使用strings.Join拼接字符串效率更高
}

func join1(arg []string) {
	defer timeCost()()	// cost 9ns
	str := ""
	for _, s := range arg {
		str = str + s + " "
	}
	fmt.Println(str)
}

func join2(arg []string) {
	defer timeCost()()	// cost 2ns
	fmt.Println(strings.Join(arg, " "))
}

func timeCost() func() {
	start := time.Now()
	return func() {
		milliseconds := time.Since(start).Microseconds()
		fmt.Printf("cost %vns\n", milliseconds)
	}
}
