package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("接通tcp失败：", err)
	}
	defer conn.Close()
	fmt.Println("客户端tcp连接成功")

	buf := make([]byte, 2048)
	go func() {
		for {
			n, _ := os.Stdin.Read(buf)
			if string(buf[:n]) == "exit\r" || string(buf[:n]) == "exit\r\n" {
				conn.Close()
				os.Exit(0)
			}
			conn.Write(buf[:n])
		}
	}()

	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("服务器已经关闭")
			break
		}
		fmt.Println("服务器回复：", string(buf[:n]))
	}
}
