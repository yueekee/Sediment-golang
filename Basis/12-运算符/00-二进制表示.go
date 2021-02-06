package main

import "fmt"

// 二进制表示法只使用数字0和1来表示数字。这个数字称之为位（bit）。
func main() {
	fmt.Printf("%3d: %08b\n", 0, 0)
	fmt.Printf("%3d: %08b\n", 1, 1)
	fmt.Printf("%3d: %08b\n", 2, 2)
	fmt.Printf("%3d: %08b\n", 3, 3)
	fmt.Printf("%3d: %08b\n", 4, 4)
	fmt.Printf("%3d: %08b\n", 5, 5)
	fmt.Printf("%3d: %08b\n", 6, 6)
	fmt.Printf("%3d: %08b\n", 7, 7)
	fmt.Printf("%3d: %08b\n", 8, 8)
	fmt.Printf("%3d: %08b\n", 16, 16)
	fmt.Printf("%3d: %08b\n", 32, 32)
	fmt.Printf("%3d: %08b\n", 64, 64)
	fmt.Printf("%3d: %08b\n", 128, 128)
}

/*
  0: 00000000
  1: 00000001
  2: 00000010
  3: 00000011
  4: 00000100
  5: 00000101
  6: 00000110
  7: 00000111
  8: 00001000
 16: 00010000
 32: 00100000
 64: 01000000
128: 10000000
*/
