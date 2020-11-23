package main

import "fmt"

func main() {
	sli := []int{0, 1}
	sli2 := []int{}
	for i, v := range sli {
		fmt.Printf("sli[%v]=%v\n", i, v)
	}
	fmt.Println("-------------")
	for i, v := range sli2 {
		fmt.Printf("sli[%v]=%v\n", i, v)
	}
}
