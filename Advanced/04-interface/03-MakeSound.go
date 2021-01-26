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

func main() {
	var toy NoiseMaker
	toy = Whistle("111") // 将一个满足NoiseMaker的类型的值赋给变量
	toy.MakeSound()      // Tweet!

	toy = Horn("222")
	toy.MakeSound() // Honk!

	// -------------- 或者可以将行数的参数定义为接口类型
	play(Whistle("234")) // Tweet!
	play(Horn("345"))    // Honk!
}

func play(n NoiseMaker) {
	n.MakeSound()
}
