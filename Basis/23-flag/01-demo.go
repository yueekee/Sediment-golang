package main

import (
	"flag"
	"fmt"
)

/*flag总结
实现命令行参数的解析
*/

// 定义命令行参数对应的变量，这两个变量都是指针类型
var cliName = flag.String("name", "nick", "Input Your Name")
var cliAge = flag.Int("age", 28, "Input Your Age")

var gender string

func main() {
	// 初始化变量 gender
	flag.StringVar(&gender, "gender", "male", "put your gender")
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	// flag.Args() 函数返回没有被解析的命令行参数
	// func NArg() 函数返回没有被解析的命令行参数的个数
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	// 输出命令行参数
	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", gender)
}

/*
使用命令：go run 01-demo.go -name=eric -age=18 -gender=female help
args=[help], num=1
arg[0]=help
name= eric
age= 18
gender= female
*/
