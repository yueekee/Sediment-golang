package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func main() {
	//mapSafe()
	mapSafe2()
}

func mapSafe() {
	m := make(map[int]int)
	go func() { //开一个协程写map
		for i := 0; i < 10000; i++ {
			lock.Lock() //加锁
			m[i] = i
			lock.Unlock() //解锁
		}
	}()
	go func() { //开一个协程读map
		for i := 0; i < 10000; i++ {
			lock.Lock() //加锁
			fmt.Println(m[i])
			lock.Unlock() //解锁
		}
	}()
	time.Sleep(time.Second * 20)

	lock.Lock()
	for i := 0; i < 10000; i++ {
		fmt.Println(m[i])
	}
	lock.Unlock()
}

// 使用这种加锁解锁也可以，且效率更高
func mapSafe2() {
	m := make(map[int]int)
	go func() { //开一个协程写map
		lock.Lock()
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
		lock.Unlock()
	}()
	go func() { //开一个协程读map
		lock.Lock()
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
		lock.Unlock()
	}()
	time.Sleep(time.Second * 20)
}
