package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := measureResponeTime(a)
	bDuration := measureResponeTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponeTime(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}