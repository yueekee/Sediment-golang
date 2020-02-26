package main

import "fmt"

// 将第i位的数字和前面i-1位队列中的数字比较，如果队列中的数对第i位数大，将队列中的数后移。

func main() {
	num := []int{20, 18, 25, 36, 15, 5}
	var i, j int
	for i = 1; i < len(num); i++ {
		temp := num[i]
		for j = i; j > 0 && num[j-1] > temp; j-- {
			num[j] = num[j-1]
		}
		num[j] = temp
	}
	fmt.Println(num)
}
