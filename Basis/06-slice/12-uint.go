package main

import (
	"fmt"
)

/*总结：
uint中 -1 表示 64位/32位的最大数-1的值
*/

func main() {
	var a uint = 1
	var b uint = 2
	fmt.Println(a-b)	// 18446744073709551615
	// 1<<64 - 1 = 18446744073709551615
}
