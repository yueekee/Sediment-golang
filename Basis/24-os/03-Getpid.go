package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid())	// 进程ID
	fmt.Println(os.Getppid())	// 父进程ID
	//pipe, writer, err := os.Pipe() // 创建管道
}
