package main

import "fmt"

// 接口即约定

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) { // *ByteCounter 满足io.writer接口的约定
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "eric"
	fmt.Fprintf(&c, "hello, %s", name) // 这里可以直接使用c作为io.Writer
	fmt.Println(c)
}
