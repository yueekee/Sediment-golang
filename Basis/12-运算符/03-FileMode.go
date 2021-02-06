package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.O_RDONLY, os.O_WRONLY, os.O_RDWR, os.O_CREATE, os.O_APPEND)
	// 0 1 2 512 8
}
