package main

import "fmt"

/* 总结：
`*` 在声明时表示为指针类型，伴随变量时表示取指针空间内的值
`&` 表示取指针空间的地址
*/

func main() {
	num := 10
	var p *int	// 声明p为int指针类型
	p = &num	// p里存放的num（int类型）的地址
	fmt.Println("p:", p)   // p: 0x1100a0b0
	fmt.Println("*p:", *p) // *p: 10

	*p = 100
	fmt.Println("p:", p)   // p: 0x1100a0b0，地址没有改变
	fmt.Println("*p:", *p) // *p: 100，只是值改变了
	fmt.Println("num:", num) // num: 100，值也跟着改变了
}
