package main

import "fmt"

type Cat struct {
	Animal
	FeatureA string
}

type Dog struct {
	Animal
	FeatureB string
}

func (cat *Cat) Run() {
	fmt.Println("cat run")
}

func (dog *Dog) Run() {
	fmt.Println("dog run")
}

func main2() {
	p := NewAnimal()
	p.SetName("miao")

	dog := Dog{Animal: *p}
	fmt.Println(dog.GetName())
}