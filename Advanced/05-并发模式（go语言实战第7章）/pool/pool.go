// Package Pool create.
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

/*
使用有缓冲的通道实现资源池，来管理可以在任意数量的goroutine之间共享及独立使用的资源
*/

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer            // 作为一个有缓冲的通道，用来保存共享的资源。
	factory   func() (io.Closer, error) // 当池需要一个新资源时，可以由这个函数来创建
	closed    bool
}

var ErrPoolClosed = errors.New("pool has been closed")

// NewPool return *Pool.
func NewPool(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 检查是否有空闲的资源
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	// 因为没有空闲资源可用，所以提供一个新资源
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	// 如果池已经被关闭，销毁这个资源
	if p.closed {
		r.Close()
		return
	}

	select {
	// 向channel写入数据，即将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")

	// 如果队列已满，则关闭这个资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// 让资源池停止工作，并关闭所有现有的资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// 关闭通道。关闭资源前如果不关闭通道会发生死锁。
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}
