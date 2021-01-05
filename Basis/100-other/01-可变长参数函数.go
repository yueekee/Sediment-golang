package main

import "fmt"

func main() {
	severalInts(1)       // [1]
	severalInts(1, 2, 3) // [1 2 3]
}

// 这里numbers变量被存储在作为参数的切片中。
func severalInts(numbers ...int) {
	fmt.Println(numbers)
}
