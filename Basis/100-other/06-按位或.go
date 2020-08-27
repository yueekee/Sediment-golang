package main

import "fmt"

func main() {
	status := 1	// 0001
	num := 2 // 0010  3-> 0011
	if true {
		num |= status
	}
	fmt.Println("num:", num) // 3
}