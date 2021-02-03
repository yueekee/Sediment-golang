package main

import "fmt"

// 打印从start到end的所有值
func count(start, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
	fmt.Printf("Returning from count(%d, %d)\n", start, end)
}

func main() {
	count(1, 3)
}

/*
count(1, 3) called
1
count(2, 3) called
2
count(3, 3) called
3
Returning from count(3, 3)
Returning from count(2, 3)
Returning from count(1, 3)
*/