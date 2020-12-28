package main

import (
	"fmt"
	"time"
)

func main() {
	var m = map[string]string{"1": "hello"}

	for k, v := range m {
		k, v := k, v // copy for capturing by the goroutine
		go func() {
			fmt.Println(k, v)
		}()
	}

	for k, v := range m {
		x := struct{ k, v string }{k, v} // copy for capturing by the goroutine
		go func() {
			fmt.Println(x.k, x.v)
		}()
	}

	time.Sleep(time.Second)
}

/*
x := struct{ k, v string }{k, v} 使用临时结构体可以将局部变量升级为堆分配。
编译器通常无法证明多个变量具有相同的生存时间，因此它会分别分配每个此类变量。使用临时结构体可以避免这个问题。
但是这种优化影响可读性，所以应合理地使用。
*/