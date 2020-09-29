package main

import (
	"fmt"
	"time"
)

/* 总结：
channel要和go程搭配使用，不然会deadlock
读端的channel会等待写端的channel写数据
*/

func main() {
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
		go func(i int) {
			fmt.Printf("%d, 子go程\n", i)
		}(i)
	}

	time.Sleep(100*time.Millisecond)
	for i := 0; i < 10; i++ {
		<-ch
	}

	fmt.Println("主go程")

	fmt.Println("主go程11")
}
