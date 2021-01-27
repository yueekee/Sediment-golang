package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 1*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second): // 1s超时走这里
		fmt.Println("oversleep")
	case <-ctx.Done(): // 到了超时时间走这里
		fmt.Println(ctx.Err()) // context deadline exceeded
	}
}
