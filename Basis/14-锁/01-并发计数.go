package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

// 这里每次打印的count值不一样，因为count++不是原子操作
// 使用go run -race 01-mutex.go 命令，可以查看哪里发生了并发问题
