package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

// 实现接口

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer) // *bytes.Buffer有Write方法
	w = time.Second       // time.Duration缺少Write方法

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	rwc = new(bytes.Buffer) // *bytes.Buffer缺少Close方法

	rwc = w // 右侧表达式也是一个接口，io.Writer缺少Close方法
	w = rwc // io.ReadWriteCloser有write方法
}
