package main

import "fmt"

func GetValue() int {
	return 1
}

func main() {
	//i := GetValue()
	var i interface{}	// 改为interface{}
	switch i.(type) {	// 报错：i不是interface{}
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("unknown")
	}
}