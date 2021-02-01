package main

import "fmt"

type Whistle string
func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string
func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface { // 代表任何含有MakeSound方法的类型
	MakeSound()
}

type Robot string
func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func main() {
	var toy NoiseMaker
	toy = Whistle("111") // 将一个满足NoiseMaker的类型的值赋给变量
	toy.MakeSound()      // Tweet!

	toy = Horn("222")
	toy.MakeSound() // Honk!

	// --------------
	// 或者可以将行数的参数定义为接口类型
	play(Whistle("234")) // Tweet!
	play(Horn("345"))    // Honk!

	// --------------
	var noiseMaker NoiseMaker = Robot("333")	// 定义一个接口类型，并且将满足接口的类型值赋非它
	noiseMaker.MakeSound()	// 调用接口的方法
	var robot Robot = noiseMaker.(Robot)	// 使用类型断言取回具体类型
	robot.Walk()	// 调用在具体类型（而不是接口）上定义的方法。
}

func play(n NoiseMaker) {
	n.MakeSound()
}
