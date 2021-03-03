package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	ch := make(chan int)
	// 起5个go程写数据，用主go程读数据
	for i := 0; i < 5; i++ {
		go func() {
			ch <- rand.Intn(100)
		}()
		fmt.Println(<-ch)
	}
}

// 写法1
func f1() {
	rand.Seed(time.Now().Unix())
	ch := make(chan int)
	// 起一个go程，写入5个数据
	go func() {
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(10)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
