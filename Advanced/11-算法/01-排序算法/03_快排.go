package main

import "fmt"

/*
快速排序的原理，首先找到一个数pivot把数组‘平均’分成两组，使其中一组的所有数字均大于另一组中的数字，
此时pivot在数组中的位置就是它正确的位置。然后，对这两组数组再次进行这种操作。
*/
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
	//startTime := time.Now().UnixNano()
	//for i := 0; i < 100000; i++ {
	//	arr := []int{12, 8, 38, 65, 2, 93}
	//	quickSort(arr, 0, len(arr)-1)
	//}
	//endTime := time.Now().UnixNano()
	//
	//nanoSeconds:= float64(endTime - startTime)
	//
	////fmt.Println(arr)
	//fmt.Printf("second:%f", nanoSeconds/100000)

	sli := []int{12, 8, 38, 65, 2, 93}
	quickSort(sli, 0, len(sli)-1)
	fmt.Println(sli)	// [2 8 12 38 65 93]
}
