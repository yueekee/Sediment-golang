package main

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

/* 
单线程版基准测试时间为2.23s
go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/liankui/Sediment-golang/GoWithTests/10-concurrency
BenchmarkCheckWebsite-16               1        2236162528 ns/op
PASS
ok      github.com/liankui/Sediment-golang/GoWithTests/10-concurrency   3.473s

并发版的基准测试时间为0.023s，快了100倍
 go test -bench=.                    
goos: darwin
goarch: amd64
pkg: github.com/liankui/Sediment-golang/GoWithTests/10-concurrency
BenchmarkCheckWebsite-16              56          23228837 ns/op
PASS
ok      github.com/liankui/Sediment-golang/GoWithTests/10-concurrency   2.265s
 */