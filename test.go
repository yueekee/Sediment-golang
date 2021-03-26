package main

import "fmt"

func test(x byte)  {
	fmt.Println(x)
}

func main() {
	var a byte = 0x11
	fmt.Println(a)
	var b uint8 = a
	var c uint8 = a + b
	test(c)
}