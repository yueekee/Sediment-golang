package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

var rawFileList []*os.File

func main() {
	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println(err)
		return
	}

	var connList []net.Conn
	for {
		conn, err := l.Accept()
		connList = append(connList, conn)
		if err != nil {
			fmt.Println(err)
			return
		}

		go func() {
			f, err := conn.(*net.TCPConn).File()
			if err != nil {
				fmt.Println(err)
				return
			}

			rawFile := os.NewFile(f.Fd(), "")
			rawFileList = append(rawFileList, rawFile)
			_ = rawFile
			for {
				var buf = make([]byte, 1024)
				conn.Read(buf)
				conn.Write([]byte(`HTTP/1.1 200 OK
Connection: Keep-Alive
Content-Length: 0
Content-Type: text/html
Server: Apache

`))
				runtime.GC()
			}
		}()
	}
}
