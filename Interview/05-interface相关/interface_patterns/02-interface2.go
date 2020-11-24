package main

import "fmt"

// 使用嵌套结构体的方式，减少多态的实现代码。
// 但是初始化的时候有些复杂
func main() {
	c1 := Country2{WithName{"Singapore"}}
	c2 := City2{WithName{"Singapore City"}}

	c1.PrintStr2()
	c2.PrintStr2()
}

type WithName struct {
	Name string
}

type Country2 struct {
	WithName
}

type City2 struct {
	WithName
}

type Printable2 interface {
	PrintStr2()
}

func (w WithName) PrintStr2() {
	fmt.Println(w.Name)
}

