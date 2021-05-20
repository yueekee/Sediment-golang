package main

type Animal struct {
	name string
}

func NewAnimal() *Animal {
	return &Animal{}
}

func (a *Animal) SetName(name string) {
	a.name = name
}

func (a *Animal) GetName() string {
	return a.name
}
