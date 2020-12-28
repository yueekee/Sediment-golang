package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G#1 r#1cks"
	replacer := strings.NewReplacer("#1", "o", "s","o")	// 替换是成对出现的。一对中前一个为老字符串，后一个为新字符串
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)	// Go rocko

	broken = "G# r#cks"
	replace := strings.Replace(broken, "#", "o", 2)	// n 表示替换的次数
	fmt.Println(replace)	// Go rocks
}
