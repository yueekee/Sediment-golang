package main

import (
	"os"
	"text/template"
)

//func check(err error) {
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// 在模版中插入数据，可以在模版文本中添加action操作(action用花括号{{}}表示)
// 在花括号中，指定要插入的数据或要模版执行的操作。
// 每当模版遇到action时，它将计算其内容，并在action的位置将结果插入模版文件中。
// 在一个操作中，可以使用一个点(.)的Execute方法来引用传递给它的数据值。
func main() {
	text := "------------\nAction: {{.}}\n"
	tmpl, err := template.New("test").Parse(text)
	check(err)

	err = tmpl.Execute(os.Stdout, "abc")
	check(err)
	err = tmpl.Execute(os.Stdout, 123)
	check(err)
	err = tmpl.Execute(os.Stdout, true)
	check(err)
}

/*
------------
Action: abc
------------
Action: 123
------------
Action: true
*/
