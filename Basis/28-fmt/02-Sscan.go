package main

import "fmt"

func main() {
	s, t := "test123", ""
	fmt.Sscan(s, &t)
	fmt.Println("s:", s)
	fmt.Println("t:", t)	// t: test123 将s的内容传给t

	fmt.Sscanln(s, &t)
	fmt.Println("s:", s)
	fmt.Println("t:", t)	// t: test123 将s的内容传给t

	_, err := fmt.Sscanf(s, "test%s", &t)
	fmt.Println("err:", err)
	fmt.Println("s:", s)	// s: test123
	fmt.Println("t:", t)	// t: 123 将t从s中去掉test后提取出来
}
