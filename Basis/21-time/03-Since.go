package main

import (
	"fmt"
	"time"
)

func main() {
	defer TimeCost()()
	time.Sleep(1* time.Second)
}

func TimeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start).Milliseconds()
		fmt.Printf("time cost = %v\n", tc)	// time cost = 1005
	}
}
