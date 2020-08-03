package main

import (
	"fmt"
	"strings"
)

func main() {
	test1()
	fmt.Println("----------------")
	test2()
	fmt.Println("----------------")
	test3()
	fmt.Println("----------------")
	test4()
	fmt.Println("----------------")
	test5()
	fmt.Println("----------------")
	test6()
}

func test1() {
	Url := "https://www.baidu.com/material/list"

	srcObjectName := strings.TrimLeft(Url, "https://www.baidu.com")
	fmt.Println("srcObjectName:", srcObjectName)
}

func test2() {
	Url := "https://www.baidu.com/material/list"

	srcObjectName := strings.Trim(Url, "https://www.baidu.com")
	fmt.Println("srcObjectName2:", srcObjectName)
}

func test3() {
	Url := "www.baidu.com/material/list"

	srcObjectName := strings.Trim(Url, "www.baidu.com")
	fmt.Println("srcObjectName3:", srcObjectName)
}

func test4() {
	Url := "https://www.baidu.com/material/list"

	srcObjectName := strings.TrimPrefix(Url, "https://www.baidu.com")
	fmt.Println("srcObjectName4:", srcObjectName)
}

func test5() {
	var s = "Goodbye,, world!"
	s = strings.TrimPrefix(s, ", world!")
	fmt.Println(s)
}

func test6() {
	var s = "Goodbye,, world!"
	s = strings.TrimSuffix(s, ", world!")
	fmt.Println(s)
}
