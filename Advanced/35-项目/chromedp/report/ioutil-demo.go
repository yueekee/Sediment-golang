package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("my.txt")
	defer f.Close()

	if err != nil {
		log.Println("err1:", err)
	}

	reader := bufio.NewReader(f)
	var chunk []byte
	buf := make([]byte, 1024)
	for {
		buf, err = reader.ReadBytes('\n') //读取到\n结束
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
		}
		chunk = append(chunk, buf...)
	}

	fmt.Printf("chunk:%#v\n", string(chunk))
	fmt.Println(len(chunk))
}
