package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"net/http"
)

// 测试 socket: too many open files 的问题

func main() {
	ch := make(chan int)
	num := 10000000
	go func() {
		for i := 0; i < num; i++ {
			ch <- i
		}
	}()
	for i := 0; i < num; i++ {
		go client()
	}

}

func client() {
	url := fmt.Sprintf("http://localhost:5555/id")
	fmt.Println(url)

	cli := gentleman.New()
	cli.URL(url)
	req := cli.Request()
	req.Method(http.MethodPost)
	res, err := req.Send()
	defer res.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	if !res.Ok {
		err := fmt.Errorf("Invalid server response: %d\n", res.StatusCode)
		fmt.Println(err)
		return
	}
	//fmt.Printf("res:%#v\n", res)
}
