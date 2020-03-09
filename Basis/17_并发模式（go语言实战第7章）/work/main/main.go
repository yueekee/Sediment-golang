package main

import (
	"github.com/yueekee/Sediment-golang/Basis/17_并发模式（go语言实战第7章）/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name 	string
}

// Task 实现Worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(100 * time.Millisecond)
}

func main() {
	// 创建工作池，其中包含2个执行任务得goroutine
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	// names切片里的每个名字都会创建100个goroutine来提交任务，
	// 这样一来就有一堆goroutine互相竞争，将任务提交到池里。
	for i := 0; i < 10; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()

	// 让所有工作池停止工作，等待所有现在得工作完成
	p.Shutdown()
}
