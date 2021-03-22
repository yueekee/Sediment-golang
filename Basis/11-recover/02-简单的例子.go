package main

import (
	"log"
	"time"
)

/*总结：
recover()在panic前使用，recover()和defer一起使用才生效
*/

func main() {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				log.Printf("recover: %v", e)
			}
		}()
		panic("this is panic")
	}()

	log.Println("hello world")
	time.Sleep(time.Second)
}
