package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

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
