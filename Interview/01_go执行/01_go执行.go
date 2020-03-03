package main

import (
	"fmt"
	"runtime"
	"sync"
)

/* 现象：
A:全是10，输出完全看goroutine执行时i的值是多少
B:一定输出0-9，顺序不定

第一个go func（）中i是外部for的一个变量，地址不变化，但是值都在改变。
第二个go func（）中i是函数参数，与外部for中的i完全是两个变量。
尾部(i)将发生值拷贝，go func内部指向值拷贝地址。
*/
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
