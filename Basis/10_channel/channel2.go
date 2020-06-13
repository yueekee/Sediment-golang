package main

import "fmt"

func main() {
	ch := make(chan int)
	flag := make(chan bool)

	go func() {
		for i := 1; i <= 50; i++ {
			ch <- i
			fmt.Println("写入ch：", i)
			//02_flag <- true
		}
	}()

	for i := 1; i <= 50; i++ {
		num := <-ch
		fmt.Println("读取ch：", num)
		<-flag
	}
}
