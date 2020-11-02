package main

import (
	"runtime"
	"sync/atomic"
)

func main() {
	
}

type Mutex struct {
	state int32
	sema  uint32
}


const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexWaiterShift = iota
)


func (m *Mutex) Lock() {
	// Fast path: 幸运case，能够直接获取到锁
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	awoke := false
	for {
		old := m.state
		new := old | mutexLocked // 新状态加锁
		if old&mutexLocked != 0 {
			new = old + 1<<mutexWaiterShift //等待者数量加一
		}
		if awoke {
			// goroutine是被唤醒的，
			// 新状态清除唤醒标志
			new &^= mutexWoken
		}
		if atomic.CompareAndSwapInt32(&m.state, old, new) {//设置新状态
			if old&mutexLocked == 0 { // 锁原状态未加锁
				break
			}
			runtime.Semacquire(&m.sema) // 请求信号量
			awoke = true
		}
	}
}