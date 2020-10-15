package main

import (
	"fmt"
	"net"
)

// 实现一台服务器对应多个客户端的并发操作！

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("监听失败：", err)
	}
	defer listener.Close()
	fmt.Println("服务端tcp监听成功")

	for {
		conn, _ := listener.Accept()
		fmt.Println("创建了conn")
		go func() {
			defer conn.Close()
			buf := make([]byte, 1024)

			for {
				n, _ := conn.Read(buf)
				remoteAddr := conn.RemoteAddr()
				if n == 0 {
					fmt.Println("客户端",remoteAddr.String(),"已经关闭")
					break
				}
				fmt.Println(remoteAddr.String(), "发来了", string(buf[:n]))

				conn.Write([]byte("收到。。。"))
			}
		}()
	}
}
