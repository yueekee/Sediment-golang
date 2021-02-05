package main

import "fmt"

// doMath函数接收另一个函数作为参数。传入函数必须接受两个整数并返回一个float64。
func doMath(passedFunction func(int, int) float64, a, b int) {
	result := passedFunction(a, b)
	fmt.Println(result)
}

func divide(a, b int) float64 {	// 可以传递给doMath的函数
	return float64(a) / float64(b)
}
func multiply(a, b int) float64 {
	return float64(a * b)
}

func main() {
	doMath(divide, 10, 2)	// 将divide函数传递给doMath
	doMath(multiply, 11, 3)
}

/*
5
20
*/