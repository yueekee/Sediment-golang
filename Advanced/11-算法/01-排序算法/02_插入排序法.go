package main

import (
	"fmt"
)

// 将第i位的数字和前面i-1位队列中的数字比较，如果队列中的数对第i位数大，将队列中的数后移。

func insertSort(sli []int) {
	var i, j int
	for i = 1; i < len(sli); i++ {
		tmp := sli[i]
		for j = i; j > 0 && sli[j-1] > tmp; j-- {
			sli[j] = sli[j-1]
		}
		sli[j] = tmp
	}
}

func main() {
	sli := []int{12, 8, 38, 65, 2, 93}
	insertSort(sli)
	fmt.Println(sli)	// [93 65 38 12 8 2]
}
