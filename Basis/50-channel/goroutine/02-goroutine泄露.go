package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
https://mp.weixin.qq.com/s/ql01K1nOnEZpdbp--6EDYw

channel 接收了值，但是不发送的话，同样会造成阻塞。
但在实际业务场景中，一般更复杂。基本是一大堆业务逻辑里，有一个 channel 的读写操作出现了问题，自然就阻塞了。
*/

func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan struct{}
	go func() {
		ch <- struct{}{}
	}()


	time.Sleep(time.Second)
}

