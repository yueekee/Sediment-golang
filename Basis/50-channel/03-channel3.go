package main

import (
	"fmt"
)

func producer(in chan<- int) { // 生产者
	for i := 1; i <= 5; i++ {
		in <- i
		fmt.Println("生产：", i)
	}
	close(in)
}

func consumer(out <-chan int) { // 消费者
	for num := range out {
		fmt.Println("消费：", num)
	}
}

func main() {
	ch := make(chan int, 10)
	go producer(ch)
	consumer(ch)
}
