package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "asdfsa|sdf|sdf"
	trim := strings.Split(s, "|")	// 根据"｜"切割字符串为字符串切片
	fmt.Println("len:", len(trim))	// len: 3
	fmt.Println("trim:", trim)		// trim: [asdfsa sdf sdf]

	s2 := "2020102114505434/20201021141122138_1.txt"
	fmt.Println("s2:", s2[36:])
}
