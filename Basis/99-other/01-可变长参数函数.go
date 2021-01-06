package main

import "fmt"

func main() {
	severalInts(1)       // [1]
	severalInts(1, 2, 3) // [1 2 3]

	intSlice := []int{1, 2, 3}
	severalInts(intSlice...)	// 使用int切片替代可变参数

	stringSlice := []string{"s", "d"}
	mix(1, false, stringSlice...) // 使用string切片代替可变参数
}

// 这里numbers变量被存储在作为参数的切片中。
func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}