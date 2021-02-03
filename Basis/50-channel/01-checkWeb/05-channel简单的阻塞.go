package main

import "fmt"

// channel通过blocking(阻塞)——暂停当前goroutine中的所有进一步操作来实现这一点。
// 发送操作阻塞发送goroutine，直到另一个goroutine在同一channel上执行了接收操作；
// 接收操作阻塞接收goroutine，直到另一个goroutine在同一channel上执行了发送操作。

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}
func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	go def(channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
}

/*
a
d
b
e
c
f
*/
