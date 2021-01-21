package main

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int
func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}
func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}
func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}
func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
// 如果一个类型包含接口中声明的所有方法，那么它可以在任何需要接口的地方使用，而不需要更多的声明。

func main() {
	var value MyInterface
	value = MyType(5)	// MyType的值满足MyInterface，所有可以将值赋给MyInterface的变量
	value.MethodWithoutParameters()
	value.MethodWithParameter(12.3)
	fmt.Println(value.MethodWithReturnValue())
	// value 可以调用MyInterface的任何方法
}
