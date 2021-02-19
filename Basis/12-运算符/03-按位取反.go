package main

import "fmt"

// 二进制数在内存中以补码的形式存储。
// “^x”的结果为“-（x+1）”
func main() {
	fmt.Println(^5)  //	-6
	fmt.Println(^-6) // 5
	fmt.Printf("5:%b\n", 5)
	fmt.Printf("^5:%b\n", ^5)
}

