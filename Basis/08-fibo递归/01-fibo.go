package main

import "fmt"

func main() {
	var num int
	fmt.Println("请输入需要得到的fibo数字的序号")
	fmt.Scan(&num)
	fmt.Println(fibo(num))
	fmt.Println(fibo2(num))
}

func fibo(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return fibo(num-2) + fibo(num-1)
	}
}

func fibo2(num int) int {
	a, b, sum := 0, 1, 0
	for i := 0; i < num; i++ {
		sum = (a + b) % 1000000007
		a, b = b, sum
	}
	return a
}
