package main

import (
	"fmt"
)

// for range ch时要用close

func main() {
	f1()
	f2()
	f3()
}

// 写法1
func f1() {
	ch := make(chan int)
	// 起一个go程，写入5个数据
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func f2() {
	var ch = make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <-i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println("-------: ", i)
	}
}

func f3() {
	ch := make(chan int)

	// 起5个go程写数据，用主go程读数据
	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
		fmt.Println(<-ch)
	}
}

