package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
		go write(i)
	}
	for i := 1; i <= 5; i++ {
		go read(i)
	}
	for {
		;
	}
}

var num int
var rwMu sync.RWMutex

func write(i int) {
	for {
		rwMu.Lock()
		num = rand.Intn(100)
		fmt.Printf("%d个go程写%d的\n", i, num)
		rwMu.Unlock()
		time.Sleep(time.Millisecond * 300)
	}
}

func read(i int) {
	for {
		rwMu.RLock()
		fmt.Printf("\t%d个go程读%d的\n", i, num)
		rwMu.RUnlock()
		time.Sleep(time.Millisecond * 300)
	}
}


