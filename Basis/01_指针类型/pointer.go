package main

import "fmt"

/* 总结：
* 在声明时表示为指针类型，伴随变量时表示取指针空间内的值
& 表示取指针空间的地址
*/

func main() {
	num := 10
	var p *int
	p = &num
	fmt.Println("p:", p)   // p: 0x1100a0b0
	fmt.Println("*p:", *p) // *p: 10

	*p = 100
	fmt.Println("p:", p)   // p: 0x1100a0b0，地址没有改变
	fmt.Println("*p:", *p) // *p: 100，只是值改变了
}
