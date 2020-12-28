package work

import (
	"sync"
)

/*
使用无缓冲的通道来创建一个goroutine池，这些goroutine执行并控制一组动作，让其并发执行。
使用无缓冲的通道不会有工作在队列里丢失或者卡住，所有工作都会被处理。
1.知道什么时候goroutine池正在执行工作
2.如果池里的所有goroutine都忙，无法接受心得工作，可及时通过通道来通知调用者
*/

type Worker interface {
	Task()
}

type Pool struct {
	work 	chan Worker
	wg 		sync.WaitGroup
}

func New(maxGoroutine int) *Pool {
	p := Pool{
		work: 	make(chan Worker),
	}

	p.wg.Add(maxGoroutine)
	for i := 0; i < maxGoroutine; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <-w
}

// Shutdown 等待所有goroutine停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}


