package main

import "fmt"

// 具有一级函数的语言中，可以将函数赋值给变量，然后使用这些变量来调用函数。
// 当调用其它函数时，函数也可以作为参数传递。

func sayHi() {
	fmt.Println("Hi")
}

func main() {
	var myFunction func() // 声明一个类型为func()的变量，该变量可以保存一个函数
	myFunction = sayHi    // 将sayHi函数赋值给变量
	myFunction()          // 调用存储在变量中函数
}
