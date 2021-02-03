package main

import "fmt"

// recover()本身的值为nil
// 延迟调用recover()时，recover会写入panic中的值
// panic()可以接收任何类型；recover()的返回值可以是任何类型

// panic和recover类似其它语言的"exception"，但go语言本身不鼓励使用panic和recover
// exception可以使程序流程更加复杂，且只是抛出一个exception，之后并没有正确地处理它。
// 直接在函数中处理错误会使函数的代码变长，但是比根本不处理错误要好得多。

func calmDown1() {
	fmt.Println(recover())
}

func calmDown2() {
	p := recover()
	err, ok := p.(error)	// 断言panic值的类型为"error"
	if ok {
		fmt.Println(err.Error())
	}
}

func main() {
	fmt.Println(recover())	// <nil>
	fmt.Println("-----------")

	//defer calmDown1()	// oh no
	//panic("oh no") // 这是将从"recover"返回的值

	fmt.Println("-----------")

	defer calmDown2()
	err := fmt.Errorf("there's an error")
	panic(err)
}
