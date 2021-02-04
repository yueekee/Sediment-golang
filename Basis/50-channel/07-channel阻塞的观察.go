package main

import (
	"fmt"
	"time"
)

func reportNqp(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNqp("sending goroutine", 2)	// 先休眠了2s，然后写入channel值
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNqp("receiving goroutine", 5) // 休眠5s后，接收channel里的值
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
