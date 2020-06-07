package main

import (
	"fmt"
	"os"
)

func main() {
	url2 := "http://wx.qlogo.cn/mmopen/kr9ShguI1uzOvkB15VXrdPWegpCW6OHd9sbdgPmicVTIMibT3bMd226KZb4ibv1XQDjV8lH0DdpjZRQ6vkO47zMibjU8cEQWGhrD/0"
	//list := os.Args
	//if len(list) == 1 {
	//	fmt.Println("无内容")
	//	return
	//}
	//fileName := list[1]

	info, err := os.Stat(url2)//Stat获取文件属性
	if err != nil {
		fmt.Println("os.Stat err =",err)
		return
	}

	fmt.Println("name =",info.Name())
	fmt.Println("size =",info.Size())
	fmt.Println("mode =",info.Mode())
	fmt.Println("modtime =",info.ModTime())
	fmt.Println("isDir =",info.IsDir())
	fmt.Println("sys =",info.Sys())

}
