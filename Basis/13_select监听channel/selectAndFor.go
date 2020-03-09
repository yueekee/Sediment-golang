package main

import (
	"context"
	"fmt"
	"log"
	"time"
)
// 使用for和select后break跳转不出来，这里使用label进行解决

func main() {
	d := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	DONE:
	for {
		select {
		case <- time.After(1 * time.Second):
			fmt.Println("超时")
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			break DONE
		}
	}
	log.Println("done")

}
