package main

import "fmt"

/*总结：
struct可以定义任意基础类型
可以把基础类型的值转换为定义的类型，如carFuel = Gallons(10.0)
*/

type Gallons float64
type Liters float64

// 一旦定义了一个类型，你可以把任何基础类型的值转换为定义的类型。
func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(10.0)
	carFuel += ToGallons(Liters(10.0))
	busFuel += ToLiters(Gallons(10.0))

	fmt.Println(carFuel, busFuel)
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

