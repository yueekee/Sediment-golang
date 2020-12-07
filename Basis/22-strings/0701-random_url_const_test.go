package main

import (
	"math/rand"
	"testing"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

func RandStringRunes1(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandStringRunes2(n int) string {
	b := make([]rune, n)
	l := len(letterRunes)
	for i := range b {
		b[i] = letterRunes[rand.Intn(l)]
	}
	return string(b)
}

const n = 10

func BenchmarkRunes1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringRunes1(n)
	}
}

func BenchmarkRunes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringRunes2(n)
	}
}

/*
go test -run=none -bench=. -benchmem 0701-random_url_const_test.go
goos: darwin
goarch: amd64
BenchmarkRunes1-16       4239441               261 ns/op              64 B/op          2 allocs/op
BenchmarkRunes2-16       4625017               261 ns/op              64 B/op          2 allocs/op
PASS
ok      command-line-arguments  3.647s
测试的结果，len()在for循环的外面声明和在里面声明性能影响不大，不用太在意这个写法的区别
*/