package main

import (
	"fmt"
)

// 建立一个接口，各个类实现要该接口，就必须实现该接口中所有的方法
// 注意：program to interface, not to implementation（即接口只做定义，不做实现）
func main() {
	d1 := Country3{"USA"}
	d2 := City3{"Los Angeles"}

	PrintStr3(d1)
	PrintStr3(d2)
}

type Country3 struct {
	Name string
}

type City3 struct {
	Name string
}

type Stringable interface {
	ToString() string
}

func (c Country3) ToString() string {
	return "Country = " + c.Name
}

func (c City3) ToString() string {
	return "City = " + c.Name
}

func PrintStr3(p Stringable) {
	fmt.Println(p.ToString())
}
