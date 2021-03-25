package main

import (
	"fmt"
	"sync"
)

func main() {
	//wrongTest()
	rightTest()
}

func wrongTest() {
	var mu sync.Mutex
	type Person struct {
		Age int
	}
	persons := make([]Person, 3)
	for i := range persons {
		mu.Lock()
		defer mu.Unlock()	// 出现死锁错误
		persons[i].Age = i
	}
	fmt.Println(persons)
}

func rightTest() {
	var mu sync.Mutex
	type Person struct {
		Age int
	}
	persons := make([]Person, 3)
	for i := range persons {
		func() {
			mu.Lock()
			defer mu.Unlock()	// 如果要想用defer，可以使用匿名函数func(){}()
			persons[i].Age = i
		}()

	}
	fmt.Println(persons)
}