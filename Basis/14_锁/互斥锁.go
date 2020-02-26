package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go user1()
	go user2()
	for {

	}
}

var mu sync.Mutex

func printer(str string) {
	mu.Lock()
	defer mu.Unlock()
	for _, s := range str {
		fmt.Printf("%c", s)
		time.Sleep(time.Millisecond * 300)
	}
}

func user1() {
	printer("hello")
}

func user2() {
	printer("world")
}
