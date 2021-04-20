package main

import "fmt"

// todo 死锁

func main() {
	c1 := make(chan int, 0)
	c2 := make(chan int, 0)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- 1
			fmt.Println("cat")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-c1
			fmt.Println("dog")
			c2 <- 2
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-c2
			fmt.Println("fish")
		}
	}()

	select {

	}
}
