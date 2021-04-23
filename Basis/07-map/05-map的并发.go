package main

import (
	"fmt"
	"sync"
)
/*总结：
如果直接使用go程进行写读，会报错 fatal error: concurrent map read and map write
解决办法：
方案1：map的写和读操作进行加锁，可以保证map的并发安全
方案2：使用sync.Map并发安全，但会有一些性能损失
*/

func main() {
	f1()
	f2()

	for  {
		
	}
}

func f1() {
	m := make(map[int]int)
	var mu sync.Mutex

	go mapStore(m, &mu)
	go mapLoad(m, &mu)
}

func f2() {
	var smap sync.Map
	go smapStore(&smap)
	go smapLoad(&smap)
}

// map的写操作
func mapStore(m map[int]int, mu *sync.Mutex) {
	mu.Lock()
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	mu.Unlock()
}

// map的读操作
func mapLoad(m map[int]int, mu *sync.Mutex) {
	mu.Lock()
	for i := 0; i < 1000; i++ {
		fmt.Println(i, ":",m[i])
	}
	mu.Unlock()
}

func smapStore(smap *sync.Map) {
	for i := 0; i < 1000; i++ {
		smap.Store(i, i)
	}
}

func smapLoad(smap *sync.Map) {
	for i := 0; i < 1000; i++ {
		fmt.Println(smap.Load(i))
	}
}