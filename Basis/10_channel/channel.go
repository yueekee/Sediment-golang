package main

import (
	"fmt"
)

/* 总结：
channel要和go程搭配使用，不然会deadlock
读端的channel会等待写端的channel写数据
*/

func main() {
	ch := make(chan bool)
	go func() {
		fmt.Println("子go程")
		ch <- true
	}()
	fmt.Println("主go程")
	<-ch
	fmt.Println("主go程11")
}
