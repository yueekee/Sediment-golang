package main

import "fmt"

/*
foo()函数中，字段ls依旧可以看成是指向底层数组的指针。
*/

type T struct {
	ls []int
}

func foo(t T) {
	t.ls[0] = 100
}

func main() {
	var t = T{
		ls: []int{1, 2, 3},
	}

	foo(t)
	fmt.Println(t.ls[0])	// 100
}
