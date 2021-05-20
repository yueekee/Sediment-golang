package main

import "fmt"

type AnimalSounder interface {
	MakeDNA()
}

func MakeSomeDNA(animalSounder AnimalSounder) {
	animalSounder.MakeDNA()
}

func (c *Cat) MakeDNA() {
	fmt.Println("cat sound")
}

func (d *Dog) MakeDNA() {
	fmt.Println("dog sound")
}

func main() {
	MakeSomeDNA(&Cat{})
	MakeSomeDNA(&Dog{})
}

/*
三大特性：“封装、继承、多态”
五大原则 “单一职责原则（SRP）、开放封闭原则（OCP）、里氏替换原则（LSP）、依赖倒置原则（DIP）、接口隔离原则（ISP）"
*/