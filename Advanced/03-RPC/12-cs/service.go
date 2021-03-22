package main

import (
	"log"
	"net/rpc"
)

func main() {
	//rpc的与服务端建立网络连接
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		log.Println("网络连接失败")
	}

	var val int
	//远程调用函数（被调用的方法，传入的参数，返回的参数）
	err = cli.Call("Num.GetInfo", 123, &val)
	if err != nil {
		log.Println("打call失败")
	}
	log.Println("返回结果", val)
}
