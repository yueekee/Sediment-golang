package main

import (
	"testing"
)

func BenchmarkSlice(b *testing.B) {
	data := make([]int, 0)
	for i := 0; i < b.N; i++ {
		data = append(data, i)
	}
}

func BenchmarkSliceCap(b *testing.B) {
	data := make([]int, 79935294)
	for i := 0; i < b.N; i++ {
		data = append(data, i)
	}
}
