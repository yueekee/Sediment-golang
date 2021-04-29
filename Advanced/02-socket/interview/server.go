package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}

		go func() {
			defer conn.Close()
			buf := make([]byte, 1024)
			for {
				l, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err)
				}
				if l > 0 {
					fmt.Print(string(buf[:l]))
				}  else {
					break
				}
				conn.Write([]byte("server resp"))
			}
		}()
	}
}