package main

import (
	"fmt"
)

type Test struct {
	Name string
}

/*
map的value本身是不可寻址的，因为map是可以自动扩容的，扩容后的地址会发生变化
使用list["name"].Name，go程序会报错
若改为var list map[string]*Test上面的使用就没有问题
*/
var list map[string]Test

func main() {

	list = make(map[string]Test)
	name := Test{"xiaoming"}
	list["name"] = name
	//list["name"].Name = "Hello"
	fmt.Println(list["name"])
}
