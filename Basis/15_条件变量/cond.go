package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

func producer(in chan <-int) {
	for {
		cond.L.Lock()
		for len(in) == 3 {
			cond.Signal()		// 等待cond的一个线程
			cond.Wait()			// Wait自行解锁c.L并阻塞当前线程，在之后线程恢复执行时，Wait方法会在返回前锁定c.L。
			// 和其他系统不同，Wait除非被Broadcast或者Signal唤醒，不会主动返回。
		}
		num := rand.Intn(100)
		fmt.Println("生产了", num)
		in <-num
		cond.L.Unlock()
	}
}

func consumer(out <-chan int) {
	for {
		cond.L.Lock()
		for len(out) == 0 {
			cond.Signal()
			cond.Wait()
		}
		num := <-out
		fmt.Println("消费了", num)
		cond.L.Unlock()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cond.L = new(sync.Mutex)
	ch := make(chan int, 3)
	for i := 1; i <= 5; i++ {
		go producer(ch)
	}
	for j := 1; j <= 5; j++ {
		go consumer(ch)
	}
	select {}
}
