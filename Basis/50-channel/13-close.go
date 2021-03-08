package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*总结：close()
close后的channel可以读取数据
close后的channel不能再写入数据，会报错 panic: send on closed channel
*/

func main() {
	rand.Seed(time.Now().Unix())
	ch := make(chan int, 10)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		ch <- rand.Intn(10)
	//	}
	//}()
	time.Sleep(time.Second)
	close(ch)

	go func() {
		fmt.Println(<-ch)
	}()
	time.Sleep(time.Second)

	go func() {
		ch <- 10
	}()
	select {}
}
