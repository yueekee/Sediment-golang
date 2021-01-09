package main

import "fmt"

type MyType string

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("a value")
	pointer := &value
	value.method()
	value.pointerMethod()	// 值类型自动转换为指针
	pointer.method()		// 指针类型自动转换为值
	pointer.pointerMethod()

	//MyType("a value").pointerMethod()	// 报错。
	// 当调用一个指针类型的接收器时，如果接收器保存在变量中，Go会自动将值转换为指针类型。如果没有保存而直接调用则会报错。
}

//Method with value receiver
//Method with pointer receiver
//Method with value receiver
//Method with pointer receiver
