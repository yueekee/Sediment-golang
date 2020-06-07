package main

import (
	"fmt"
)

func print() bool {
	fmt.Println("+++++")
	return false
}

func main() {
	//url2 := "http://wx.qlogo.cn/mmopen/kr9ShguI1uzOvkB15VXrdPWegpCW6OHd9sbdgPmicVTIMibT3bMd226KZb4ibv1XQDjV8lH0DdpjZRQ6vkO47zMibjU8cEQWGhrD/0"
	//header2, err := http.Head(url2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("header2:%+v\n", header2.Header["X-Errno"])
	if !print() {
		fmt.Println("123")
	}
}
