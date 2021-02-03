package main

import "fmt"

// <-运算符
// 向channel上发送值：myChannel <- "hi" 从发送的值指向发送该值的channel
// 接收来自channel的值： <- myChannel  有点像从channel中取出一个值

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	receiveValue := <-myChannel
	fmt.Println(receiveValue)	// hi
}
