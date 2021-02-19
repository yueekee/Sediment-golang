package main

import "fmt"

// 使用range可以安全的遍历切片，不用担心数组越界的情况

func main() {
	sli := []int{0, 1}
	for i, v := range sli {
		fmt.Printf("sli[%v]=%v\n", i, v) // 正常
	}

	fmt.Println("-------------")

	sli2 := []int{}
	for i, v := range sli2 {
		fmt.Printf("sli[%v]=%v\n", i, v) // 正常
	}
	fmt.Println("-------------")

	for i := 0; i <= len(sli2); i++ {
		fmt.Printf("sli2[%v]=%v\n", i, sli2[i]) // panic: runtime error: index out of range [0] with length 0
	}
}
