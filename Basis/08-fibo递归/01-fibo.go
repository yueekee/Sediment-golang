package main

import "fmt"

func main() {
	var num int
	fmt.Println("请输入需要得到的fibo数字的序号")
	fmt.Scan(&num)
	n := fibo(num)
	fmt.Println(n)
}

func fibo(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return fibo(num-2) + fibo(num-1)
	}
}
