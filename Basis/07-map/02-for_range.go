package main

import (
	"fmt"
)

func main() {
	var m = make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3

	for s, i := range m {
		fmt.Println(s, i)
	}

	// 切片的forr是从0索引开始的
	var n = []int{4, 2, 3}
	for i, i2 := range n {
		fmt.Println(i, i2)
	}

	for s := range m { // 仅仅处理值。
		fmt.Println("s:", s)
	}

}
