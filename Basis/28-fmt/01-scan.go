package main

import "fmt"

func main() {
	var name string
	fmt.Printf("请输入内容:")
	fmt.Scan(&name)
	fmt.Println("name:", name)

	fmt.Scanf("%s", &name)
	fmt.Println("name:", name)
}
