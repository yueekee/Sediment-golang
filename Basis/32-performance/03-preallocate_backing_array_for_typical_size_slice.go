package main

import "fmt"

func main() {
	x := MakeX()
	fmt.Printf("x:%v\n", x) // x:&{[] [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]}
}

type X struct {
	buf      []byte
	bufArray [16]byte // Buf usually does not grow beyond 16 bytes.
}

func MakeX() *X {
	x := &X{}
	// Preinitialize buf with the backing array.
	x.buf = x.bufArray[:0]
	return x
}

/*
切片数组的预分配：
如果提前知道切片的大小，可以为其分配一个备用数组
*/