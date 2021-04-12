package main

import "fmt"

/* 选择排序
拿第n个数字和之后的所有数字进行比较，大的排到前面
 */

func selectionSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := i+1; j <= len(sli)-1; j++ {
			if sli[j] > sli[i] {
				sli[j], sli[i] = sli[i], sli[j]
				fmt.Printf("i:%d,j:%d,sli:%v\n", i, j, sli)
			}
		}
	}
}

func main() {
	sli := []int{12, 8, 38, 65, 2, 93}
	selectionSort(sli)
	fmt.Println(sli)	// [93 65 38 12 8 2]
}
