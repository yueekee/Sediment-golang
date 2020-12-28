package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkStr(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		str += "x"
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteString("x")
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString("x")
	}
}

/*
✗ go test -v -run="none" -bench=. 05-string_test.go
goos: darwin
goarch: amd64
BenchmarkStr
BenchmarkStr-16                  1000000             49863 ns/op
BenchmarkStringsBuilder
BenchmarkStringsBuilder-16      407583212                3.21 ns/op
BenchmarkBytesBuffer
BenchmarkBytesBuffer-16         220828958                5.41 ns/op
PASS
ok      command-line-arguments  53.393s
*/


