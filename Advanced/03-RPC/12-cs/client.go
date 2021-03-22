package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//- 方法是导出的
//- 方法有两个参数，都是导出类型或内建类型
//- 方法的第二个参数是指针
//- 方法只有一个error接口类型的返回值
//
//func (t *T) MethodName(argType T1, replyType *T2) error

type Num int

func (n *Num) GetInfo(argType int, replyType *int) error {

	log.Println(argType)
	*replyType = 1 + argType

	return nil
}

func main() {
	//new 一个对象
	pd := new(Num)
	//注册服务
	//Register在默认服务中注册并公布 接收服务 pd对象 的方法
	rpc.Register(pd)

	rpc.HandleHTTP()
	//建立网络监听
	ln, err := net.Listen("tcp", "127.0.0.1:10086")
	if err != nil {
		log.Println("网络连接失败")
	}

	log.Println("正在监听10086")
	//service接受侦听器l上传入的HTTP连接，
	http.Serve(ln, nil)

}
