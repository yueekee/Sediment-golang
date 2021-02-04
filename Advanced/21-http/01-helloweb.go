package main

import (
	"log"
	"net/http"
)

// http.ResponseWriter 用于更新将发送到浏览器的响应的值
// *http.Request 表示来自浏览器的请求的值
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello,web")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)             // 调用viewHandler函数来生成响应
	err := http.ListenAndServe("localhost:10101", nil) // 监听服务器请求，并响应它们
	log.Fatal(err)
}
