package main

import (
	"fmt"
	"time"
)

/* 总结：
for range 使用短变量声明(:=)的形式迭代变量，需要注意的是，变量 i、v 在每次循环体中都会被重用，而不是重新声明。
各个 goroutine 中输出的 i、v 值都是 for range 循环结束后的 i、v 最终值。可以理解为闭包引用，使用的是上下文环境的值。
*/

func main() {
	var m = []int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 1)
	/*
	2 3
	2 3
	2 3
	*/
	fix()
	/*
	0 1
	2 3
	1 2
	*/
	fix2()
	/*
	0 1
	2 3
	1 2
	*/
}

func fix() {
	var m = []int{1, 2, 3}

	for i, v := range m {
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}

	time.Sleep(time.Second * 1)
}

func fix2() {
	var m = []int{1, 2, 3}

	for i, v := range m {
		i := i
		v := v
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 1)
}
