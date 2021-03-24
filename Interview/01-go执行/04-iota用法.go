package main

import "fmt"

const (
	x = iota
	_
	y
	i = iota
	j
	z = "zz"
	k
	p = iota
)

func main() {
	fmt.Println(x, y, i, j, z, k, p)	// 0 2 3 4 zz zz 7
}
