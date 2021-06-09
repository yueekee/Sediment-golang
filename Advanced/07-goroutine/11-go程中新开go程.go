package main

import (
	"fmt"
	"time"
)

// todo
func main() {
	var ns [100][100]int
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			ns[i][j] = j
		}
	}

	time.Sleep(1 * time.Second)

	for _, nums := range ns {
		go bianli(nums)
	}
}

func bianli(nums [100]int) {
	for _, num := range nums {
		go print(num)
		//fmt.Print(num, " ")
	}
}

func print(i int) {
	fmt.Print(i, " ")
}
