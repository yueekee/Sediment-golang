package main

import "fmt"

/* 闭包：
	1.匿名函数
	2.将内部的匿名函数返回出来
	3.每个闭包都要保存自己的值
*/

func main() {
	fun := fun1()
	fmt.Println(fun())		// 1
	fmt.Println(fun())		// 2

	fun2()
	fun3()
}

func fun1() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func fun2() {
	var fn[3] func()
	for i := 0; i < 3; i++ {
		fn[i] = func() {
			fmt.Println(i)
		}
	}
	fn[0]()		// 3
	fn[1]()		// 3，这里分析，i的变量只有一个且最后一个值为3，被3个函数共同访问
}

// 作以下过更改
func fun3() {
	var fn[3] func()
	for i := 0; i < 3; i++ {
		fn[i] = fun4(i)
	}
	fn[0]()		// 0, 闭包可以保存独有的信息
	fn[1]()		// 1
	fn[2]()		// 2
}

func fun4(i int) func() {
	return func() {
		fmt.Println(i)
	}
}


