package main

import "fmt"

func main() {
	s := "1"
	fmt.Println(s)
	defer func(s string) {
		s = ""
		fmt.Println(s)
	}(s)
	fmt.Println(s)

}
