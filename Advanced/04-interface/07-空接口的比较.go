package main

import "fmt"

type A interface {
}

type B interface {
}

func main() {
	a := A(1)
	b := B("1")
	if a == b {
		fmt.Println("a = b")
	}

	c := struct {
	}{}
	d := struct {
	}{}

	if c == d {
		fmt.Println("c = d")
	}
}
