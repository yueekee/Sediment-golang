package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// ioutil.ReadDir() 获取包含目录内容的切片
// filepath.Join() 用斜杠将目录路径和文件名连接起来

// 添加一个函数，打印错误信息，并体面地退出程序
func reportPanic() {
	p := recover() // 调用recover并储存它的返回值
	if p == nil {  // 如果recover返回nil，则没有panic
		return
	}
	err, ok := p.(error) // 否则，获取底层的error值
	if ok {
		fmt.Println(err) // 然后打印出来
	} else {
		panic(p)	// 如果panic值不是error，则使用相同的值恢复panic
	}
}

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path) // 获取包含目录内容的切片
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
	defer reportPanic()	// 在调用可能引起panic的代码之前，延迟调用reportPanic函数。
	panic("some other issue")
	scanDirectory("./05-struct")
}
