package main

import "fmt"

/*冒泡算法
比较两个相邻的数字的大小，如果左侧的数大则进行交换，直至将大数移到右侧。
*/

func bubbleSort(sli []int) {
	for i := 0; i < len(sli); i++ {
		for j := 1; j < len(sli) - i; j++ {
			if sli[j] < sli[j-1] {
				sli[j], sli[j-1] = sli[j-1], sli[j]
			}
			//fmt.Printf("i:%d,j:%d,sli:%v\n", i, j, sli)
		}
	}
}

func main() {
	sli := []int{12, 8, 38, 65, 2, 93, 1}
	bubbleSort(sli)
	fmt.Println(sli)	// [1 2 8 12 38 65 93]
}