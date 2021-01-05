package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximum(12.3, 33.4, 99.0)) // 99
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	fmt.Println(max) // -Inf代表了负无穷的特殊值
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
