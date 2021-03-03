package main

import "fmt"

// 如果ch内没有值，close后读取的值为0
// 如果ch内没有值，直接读取值会报死锁的错误

func main() {
	test1()	// 0
	test2()	// fatal error: all goroutines are asleep - deadlock!
}

func test1() {
	ch := make(chan int, 10)
	close(ch)
	fmt.Println(<-ch)
}

func test2() {
	ch := make(chan int, 10)
	fmt.Println(<-ch)
}
