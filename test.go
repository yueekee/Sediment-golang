package main

import "fmt"

func main() {
	x := interface{}(nil)
	t, c := x.(interface{})
	fmt.Printf("%#v\n", x)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", c)
}