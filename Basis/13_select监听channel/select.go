	package main

import (
	"fmt"
	"time"
)

/*	select实现超时退出
	select{}可以阻塞主函数
*/

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num:", num)
			case <-time.After(time.Second * 5):
				fmt.Println("超时退出")
				quit <- true
			}
		}
	}()
	<-quit
}

func zuse() {
	go func() {
		for {
			fmt.Println("golang")
		}
	}()
	select { //阻塞主程序
	}
}
