package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// 1.构建一个简单的http-rpc服务器
// 2.改为tcp-rpc服务器
// 3.改为json-rpc服务器
//type Args struct {
//	A, B int
//}

type Math int

func (m Math) Multiple(args Args, reply *int) error {	// 这里第二个参数必须是指针
	*reply = args.A + args.B
	return nil
}

//type Quotient struct {
//	Quo, Rem int
//}

func (m Math) Divide(args Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by 0")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
func main() {
	math := new(Math)
	rpc.Register(math)
	// http-rcp
	//rpc.HandleHTTP()
	//
	//if err := http.ListenAndServe(":1234", nil); err != nil {
	//	fmt.Println(err.Error())
	//}

	// tcp-rpc
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		fmt.Println("Fatal err:", err)
		os.Exit(2)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Fatal err:", err)
		os.Exit(2)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
			continue
		}
		//rpc.ServeConn(conn)
		// json-rpc
		jsonrpc.ServeConn(conn)
	}
}
