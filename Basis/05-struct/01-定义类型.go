package main

import "fmt"

type Liters float64
type Gallons float64

// 一旦定义了一个类型，你可以把任何基础类型的值转换为定义的类型。
func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)
}

