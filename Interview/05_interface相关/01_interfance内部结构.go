package main

import (
	"fmt"
)

var in interface{}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	fmt.Println(live())
	fmt.Println(live() == nil)
	fmt.Println(in)
	fmt.Println(in == nil)
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

/* interface内部结构
1.var in interface{} 这种是为nil
2.type People interface {Show()} 这种不为nil

解释：
第2种因为type类型中还有一个itab结构，data指向nil，并不代表interface为nil
*/
