package main

import "fmt"

func main() {
	//appendTest()
	appendTest2()
}

// Append() could reallocate the memory
func appendTest() {
	a := make([]int, 32)
	fmt.Printf("&a:%p\n", a)
	b := a[1:16]
	fmt.Printf("&b:%p\n", b)
	a = append(a, 1)
	a[2] = 42

	fmt.Println(a)
	fmt.Printf("&a:%p\n", a)
	fmt.Println(b)
	fmt.Printf("&b:%p\n", b)
}

/*
&a:0xc000110000
&b:0xc000110008
[0 0 42 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
&a:0xc000114000
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
&b:0xc000110008

a的地址改变了，b的地址不同于a且未改变
*/

// Append() won't reallocate the memory if capacity is enough
func appendTest2() {
	a := make([]int, 17)	// 只更新了这个长度
	fmt.Printf("&a:%p\n", a)
	b := a[1:16]
	fmt.Printf("&b:%p\n", b)
	a = append(a, 1)
	a[2] = 42

	fmt.Println(a)
	fmt.Printf("&a:%p\n", a)
	fmt.Println(b)
	fmt.Printf("&b:%p\n", b)
}

/*
&a:0xc00011a000
&b:0xc00011a008
[0 0 42 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
&a:0xc000120000
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
&b:0xc00011a008

a和b的地址都没有变化
*/