package main

import (
	"fmt"
	"time"
)

func main() {
	n := 922337203685477580 // 2的63次方是9223372036854775808
	fmt.Println(n * 10 + 7)
	fmt.Println(n * 10 + 8)
	fmt.Println(n * 10 + 9)
	fmt.Println(n * 10 + 10)

	endTime2 := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("endTime2:", endTime2)
}
