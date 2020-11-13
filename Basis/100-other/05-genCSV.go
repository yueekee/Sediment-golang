package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.csv")//创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)// 创建一个新的写入文件流
	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	w.WriteAll(data)//写入数据
	w.Flush()
}

func GenerateCSV(filePath string, data [][]string) error {
	fp, err := os.Create(filePath) // 创建文件句柄
	if err != nil {
		log.Println("创建文件["+filePath+"]句柄失败,%v", err)
		return err
	}
	defer fp.Close()
	_, err = fp.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	if err != nil {
		log.Println(err)
		return err
	}
	w := csv.NewWriter(fp) // 创建一个新的写入文件流
	err = w.WriteAll(data)
	if err != nil {
		log.Println(err)
		return err
	}
	w.Flush()
	return err
}