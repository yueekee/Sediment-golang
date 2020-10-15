package main

import (
	"fmt"
	"time"
)

// 将第i位的数字和前面i-1位队列中的数字比较，如果队列中的数对第i位数大，将队列中的数后移。

func insertSort(arr []int) {
	var i,j int
	for i = 1; i < len(arr); i++ {
		temp := arr[i]
		for j = i; j > 0 && arr[j-1] > temp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func main() {
	startTime := time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		arr := []int{12, 8, 38, 65, 2, 93}
		insertSort(arr)
	}
	endTime := time.Now().UnixNano()

	nanoSeconds:= float64(endTime - startTime)

	//fmt.Println(arr)
	fmt.Printf("second:%f", nanoSeconds/100000)
}
