package main

import (
	"fmt"
	"time"
)

/* 拿第n个数字和之后的所有数字进行比较，大的排到前面

 */
func selectionSort(arr []int) {
	for n := 0; n < len(arr)-1; n++ {
		for i := n + 1; i <= len(arr)-1; i++ {
			if arr[i] > arr[n] {
				arr[i], arr[n] = arr[n], arr[i]
			}
		}
	}
}

func main() {
	startTime := time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		arr := []int{12, 8, 38, 65, 2, 93}
		selectionSort(arr)
	}
	endTime := time.Now().UnixNano()

	nanoSeconds:= float64(endTime - startTime)

	//fmt.Println(arr)
	fmt.Printf("second:%f", nanoSeconds/100000)
}
