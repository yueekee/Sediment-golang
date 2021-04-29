package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			conn.Write([]byte(line))
		}
	}()

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		if n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}