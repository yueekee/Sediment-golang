package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// 在命令后添加参数并计算参数们的平均值
// 如：go run 02-average.go 1.1 1.2

func main() {
	//numbers, err := datafile.GetFloats("data.txt")
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
