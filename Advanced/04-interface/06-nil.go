package main

import "fmt"

// interface{}类型称为空接口，可以接受任何类型的值

func AcceptAnything(thing interface{}) {	// 接收一个空接口作为参数，但是空接口不能调用对应的方法
	fmt.Println(thing)
	whistle, ok := thing.(Whistle)
	if ok {
		whistle.MakeSound()
	}
}

// 最好还是写一个接收特定具体类型的函数
func AcceptWhistle(whistle Whistle) {
	fmt.Println(whistle)
	whistle.MakeSound()
}

func main() {
	AcceptAnything(3.14159)
	AcceptAnything("ABC")
	AcceptAnything(true)
	AcceptAnything(Whistle("meme"))

	// ------------
	AcceptWhistle(Whistle("111"))
}
