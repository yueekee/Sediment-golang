package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

// 使用 go run client.go localhost 执行命令
func main() {
	if len(os.Args) != 2 {
		fmt.Println("os.Args err:", os.Args)
		os.Exit(1)
	}
	serverAddr := os.Args[1]

	// http-rpc
	//client, e := rpc.DialHTTP("tcp", serverAddr+":1234")
	// tcp-rpc
	//client, e := rpc.Dial("tcp", serverAddr+":1234")
	// json-rpc
	client, e := jsonrpc.Dial("tcp", serverAddr+":1234")


	if e != nil {
		log.Fatal("dial err:", e)
	}

	args := Args{11, 3}
	var replay int
	err := client.Call("Math.Multiple", args, &replay)
	if err != nil {
		log.Fatal("call mul err:", e)
	}
	fmt.Printf("Math: %d+%d=%d\n", args.A, args.B, replay)

	var quo Quotient
	err = client.Call("Math.Divide", args, &quo)	// 注意这个指针的传递，因为这个函数传的是指针类型
	if err != nil {
		log.Fatal("call quo err:", e)
	}
	fmt.Printf("Math: %d / %d = %d remider %d\n", args.A, args.B, quo.Quo, quo.Rem)
}
