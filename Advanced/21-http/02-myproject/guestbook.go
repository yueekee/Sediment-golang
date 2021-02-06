package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

// http.ResponseWriter 用于更新将发送到浏览器的响应的值
// *http.Request 表示来自浏览器的请求的值
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getString("signatures.txt")
	html, err := template.ParseFiles("view.html") // 使用view.html创建一个新模版
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	check(err)
	err = html.Execute(writer, guestbook) // 将新模版写入ResponseWriter
	check(err)
	//
	//message := []byte("Hello,web")
	//_, err = writer.Write(message) // 将该数据添加到发送到浏览器的响应中
	//check(err)
}

func getString(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) { // 如果返回一个错误说文件不存在
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err()) // 报告任何扫描错误
	return lines
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html") // 使用view.html创建一个新模版
	check(err)
	err = html.Execute(writer, nil) // 将新模版写入ResponseWriter
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")	// 访问来来自请求的表单数据
	// 将签名写入txt文件
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE // 打开文件的选项
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature) // 在文件的新行上写一个签名
	check(err)
	err = file.Close()
	check(err)
	// HTTP重定向
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler) // 调用viewHandler函数来生成响应
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:10101", nil) // 监听服务器请求，并响应它们
	log.Fatal(err)
}
