package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func BenchmarkFmt(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		fmt.Sprint(rand.Int())
	}
}

func BenchmarkStrconv(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		strconv.Itoa(rand.Int())
	}
}

/*
 ✗ go test -v -bench=. 06-fmt_strcov_test.go
goos: darwin
goarch: amd64
BenchmarkFmt
BenchmarkFmt-16          9481873               118 ns/op
BenchmarkStrconv
BenchmarkStrconv-16     19896139                55.1 ns/op
PASS
ok      command-line-arguments  2.462s
*/