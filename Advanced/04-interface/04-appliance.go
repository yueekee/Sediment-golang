package main

import (
	"fmt"
)

type Appliance interface {
	TurnOn()
}

type Fan string
func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

type CoffeePot string
func (c CoffeePot) TurnOn() {
	fmt.Println("Powering up")
}
func (c CoffeePot) Brew() {
	fmt.Println("Heating up")
}

func main() {
	var device Appliance
	device = Fan("12")
	device.TurnOn()	// Spinning
	device = CoffeePot("2")
	device.TurnOn()	// Powering up
	//device.Brew()	// 报错，只能调用接口中定义的方法
}
