package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// filepath.Join() 用斜杠将目录路径和文件名连接起来

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)	// 获取包含目录内容的切片
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name()) // 用斜杠将目录路径和文件名连接起来
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	scanDirectory("./05-struct")
}
