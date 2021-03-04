package main

import "fmt"

/*总结：结构体指针
1.初始化一个结构体指针pointer，调用结构体字段有两种方式
	(*pointer).myField	// 注意有括号
	pointer.myField		// 点运算符允许通过struct的指针来访问字段
2.大型结构体的传递使用指针
*/

type myStruct struct {
	myField int
}

func main() {
	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	//fmt.Println(*pointer.myField) // 报错，go认为myField必须是一个指针
	fmt.Println((*pointer).myField)	// 3 获取指针指向的struct的值，然后访问struct的字段
	fmt.Println(pointer.myField)	// 3 点运算符允许通过struct的指针来访问字段

	pointer.myField = 9 // 也可以通过指针来赋值给struct字段
	fmt.Println(pointer.myField)
}

// 另外要注意： 使用指针传递大型struct

