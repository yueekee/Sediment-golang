package main

import (
	"context"
	"fmt"
	"time"
)

/*
type Context interface {
	// 返回context被取消的时间
	Deadline() (deadline time.Time, ok bool)

	// 返回一个channel，这个channel会在当前工作完成或者上下文被取消后关闭
	Done() <-chan struct{}

	// 返回context结束的原因，被取消：“Canceled”；超时：“DeadlineExceeded”
	Err() error

	// 获取context中key对应的值
	Value(key interface{}) interface{}
}
*/

// 为什么叫同步信号呢？
func main() {
	// 创建一个过期时间为1s的上下文
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	// 向上下文传入handle函数，该函数500ms处理传入的请求
	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("Process request with", duration)
	}
}

/*
Process request with 500ms
main context deadline exceeded
*/