package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// filepath.Join()

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)	// 获取包含目录内容的切片
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name()) // 用斜杠
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	err := scanDirectory("./../05-struct")
	if err != nil {
		log.Fatal(err)
	}
}
