package main

import (
	"fmt"
	"io"
	"net/http"
)

// 依赖注入 dependency injection

/*
io.Writer 是一个很好的通用接口，用于「将数据放在某个地方」
bytes 包中的 buffer 类型实现了 Writer 接口。
*/

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
