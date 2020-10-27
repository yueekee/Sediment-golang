package main

import "fmt"

// err != nil {}里的err会被重新赋值

func main() {
	err1 := fmt.Errorf("这是err1")
	if  err1 != nil {
		fmt.Println("----")
		err1 = nil
	}
	fmt.Println(err1)
}
