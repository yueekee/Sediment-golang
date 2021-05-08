package main

import (
	"fmt"
	"net/http"
	"time"
)

// 版本1
// func Racer(a, b string) string {
// 	aDuration := measureResponeTime(a)
// 	bDuration := measureResponeTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponeTime(url string) time.Duration {
// 	startA := time.Now()
// 	http.Get(url)
// 	return time.Since(startA)
// }

// 版本2: 使用select改进
func Racer(a, b string) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <- time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
