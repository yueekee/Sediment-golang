package main

import (
	"fmt"
	"time"
)

// 每个Go程序的main函数都是使用goroutine启动的。
// 因此每个Go程序至少运行一个goroutine。
// go语句不能使用返回值，因为在尝试使用它之前，不能保证返回值已经准备好。

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("----end main()")
}
