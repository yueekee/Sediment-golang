package main

import "fmt"

// 具有一级函数的编程语言还允许将函数作为参数传递给其它函数。

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {	// 接收另一个函数作为参数
	theFunction()
	theFunction()
}

func main() {
	twice(sayBye)	// 将sayBye函数传递给twice函数
}

/*
Bye
Bye
*/