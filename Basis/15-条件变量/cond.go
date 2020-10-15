package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

func producer(in chan<- int) {
	for {
		cond.L.Lock()
		for len(in) == 3 {	// 循环判断缓存是否已满
			cond.Signal()	// 唤醒之前的阻塞
			cond.Wait() // 阻塞-解锁-加锁
		}
		num := rand.Intn(100)	// 写入随机的数字
		in <- num
		fmt.Println("生产了", num)
		cond.L.Unlock()
	}
}

func consumer(out <-chan int) {
	for {
		cond.L.Lock()
		for len(out) == 0 {	// 循环判断缓存是否为空
			cond.Signal()
			cond.Wait()
		}
		num := <- out
		fmt.Println("消费了", num)
		cond.L.Unlock()
	}
}

func main() {
	ch := make(chan int, 3)
	cond.L = new(sync.Mutex)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		go producer(ch)
	}
	for i := 0; i < 5; i++ {
		go consumer(ch)
	}

	select {}	// 阻塞主函数
}
