package main

import (
	"log"
	"net/http"
)

//func check(err error) {
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// http.ResponseWriter 用于更新将发送到浏览器的响应的值
// *http.Request 表示来自浏览器的请求的值
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello,web")
	_, err := writer.Write(message) // 将该数据添加到发送到浏览器的响应中
	check(err)
}

func main() {
	http.HandleFunc("/hello", viewHandler)             // 调用viewHandler函数来生成响应
	err := http.ListenAndServe("localhost:10101", nil) // 监听服务器请求，并响应它们
	log.Fatal(err)
}
