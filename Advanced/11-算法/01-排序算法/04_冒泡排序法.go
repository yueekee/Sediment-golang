package main

import (
	"fmt"
	"time"
)

func bubbleSort(arr []int) {
	for i := 1; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i; j++ {
			if arr[j+1] > arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func main() {
	startTime := time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		arr := []int{12, 8, 38, 65, 2, 93}
		bubbleSort(arr)
	}
	endTime := time.Now().UnixNano()

	nanoSeconds:= float64(endTime - startTime)

	//fmt.Println(arr)
	fmt.Printf("second:%f", nanoSeconds/100000)
}