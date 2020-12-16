package main

import (
	"fmt"
	"strings"
)

// join 如果只有一个元素，第一个元素后面没有","
func main() {
	cID := []string{"1"}
	join1 := strings.Join(cID, ",")
	fmt.Println("join1:", join1)	// join1: 1

	cID2 := []string{"1", "2"}
	join2 := strings.Join(cID2, ",")
	fmt.Println("join2:", join2)	// join2: 1,2
}
