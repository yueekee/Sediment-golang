package main

import "fmt"


func main() {
	info := "史蒂夫史蒂夫1"
	infoRune := []rune(info)
	s := string(infoRune[:4])
	fmt.Println(s)	// 史蒂夫史

	// 方法二
	// 一个中文字符对应3个字节
	info = "史蒂夫史蒂夫1"
	s = info[:4*3]
	fmt.Println(s)	// 史蒂夫史
}
