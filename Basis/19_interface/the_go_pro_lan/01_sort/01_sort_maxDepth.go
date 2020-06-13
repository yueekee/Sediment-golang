package main

import "fmt"

func main() {
	//01_sort.Sort()
	depth := maxDepth(16)
	fmt.Println("maxDepth:", depth)
}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 { // >>= 右移后赋值
		fmt.Println("i:", i)
		depth++
		fmt.Println("depth", depth)
	}
	return depth * 2
}
