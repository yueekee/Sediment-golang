package main

import (
	"fmt"
	"time"
)

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		//fmt.Printf("---i=%d,j=%d\n", i, j)
		key := arr[(start + end)/2]
		//fmt.Println("\tkey=", key)
		for i <= j {
			for arr[i] < key {			// 拿到arr[i]比key大
				i++
			}
			//fmt.Println("\ti=", i)
			for arr[j] > key {			// 拿到arr[j]比key小
				j--
			}
			//fmt.Println("\tj=", j)
			if i <= j {					// 将上面两个数进行交换
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
			//fmt.Println("\tarr=", arr)
		}
		//fmt.Println("arr=", arr)
		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

func main() {
	startTime := time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		arr := []int{12, 8, 38, 65, 2, 93}
		quickSort(arr, 0, len(arr)-1)
	}
	endTime := time.Now().UnixNano()

	nanoSeconds:= float64(endTime - startTime)

	//fmt.Println(arr)
	fmt.Printf("second:%f", nanoSeconds/100000)
}
