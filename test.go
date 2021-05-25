package main

import "fmt"

var o = fmt.Print

func main1() {
	var m map[int]bool // nil
	_ = m[123]
	fmt.Printf("%#v\n", m)
	var p *[5]string // nil
	for range p {
		_ = len(p)
	}
	fmt.Printf("%#v\n", p)
	var s []int // nil
	_ = s[:]
	s, s[0] = []int{1, 2}, 9
}

func main() {
	fmt.Println(1e4)
}