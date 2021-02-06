package main

import "fmt"

func main() {
	fmt.Printf("false || false == %t\n", false || false)
	fmt.Printf("true || false == %t\n", true || false)
	fmt.Printf("true || true == %t\n", true || true)

	fmt.Printf("%b | %b == %b\n", 0, 0, 0|0) // 0 | 0 == 0
	fmt.Printf("%b | %b == %b\n", 0, 1, 0|1) // 0 | 1 == 0
	fmt.Printf("%b | %b == %b\n", 1, 1, 1|1) // 1 | 1 == 1

	// 如果左边数字中相同位置的位和右边数字中相同位置的位都是1，&运算符才能将结果中的位设置为1
	fmt.Printf("%08b\n", 170)    // 10101010
	fmt.Printf("%08b\n", 15)     // 00001111
	fmt.Printf("%08b\n", 170|15) // 10101111
}
