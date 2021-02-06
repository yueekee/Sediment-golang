package main

import (
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.OpenFile("my.txt", os.O_WRONLY, os.FileMode(0600))	// 这样将新文本插入到文件的开头，并覆盖了原来的数据
	check(err)
	_, err = file.Write([]byte("amazing!\n"))
	check(err)
	err = file.Close() 	// 未使用defer，是因为这里的错误需要处理
	check(err)
}
