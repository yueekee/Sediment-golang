package main

import "fmt"

/*
map的键可以是指针类型
指针类型其实是int64或int32类型
*/

func main() {
	m := make(map[*int]int)
	var p *int
	m[p] = 11

	fmt.Printf("%#v\n", m)	// map[*int]int{(*int)(nil):11}
	fmt.Printf("%T\n", p)	// *int
}
