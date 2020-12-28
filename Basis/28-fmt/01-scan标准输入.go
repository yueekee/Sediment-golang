package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var name string
	fmt.Print("请输入内容:")
	fmt.Scan(&name)
	fmt.Println("name:", name)

	fmt.Print("请输入内容2:")
	fmt.Scanf("%s", &name)
	fmt.Println("name:", name)

	// 方法二：使用os.Stdin接收标准输入
	fmt.Print("请输入内容3:")
	reader := bufio.NewReader(os.Stdin)	// 设置从键盘获取文本的"缓冲读取器"
	readString, err := reader.ReadString('\n')	// 返回用户输入的所有内容，知道按下<Enter>键为止
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(readString)
}
