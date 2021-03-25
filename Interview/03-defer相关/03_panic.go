package main

import "fmt"

/*
panic仅有最后一个可以被recover捕获
panic("panic")先执行，但是defer还有一个panic("defer panic")，会覆盖之前的panic("panic")
*/

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}
