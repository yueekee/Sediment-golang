package main

import "fmt"

/*
切片是一种指针类型，使用函数调用切片，可以直接改变其值
*/

func main() {
	ids := []int{1,2}
	fmt.Println("ids1:", ids)	// ids1: [1 2]

	Update(ids)
	fmt.Println("ids2:", ids)	// ids2: [1 100]
}

func Update(ids []int) {
	ids[1] = 100
}


