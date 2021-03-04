package main

import "fmt"

/*总结：接口和指针
实现接口里的类型如果是指针类型，声明时必须按指针类型声明
*/

type Toggleable interface{
	toggle()
}

type Switch string
func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

func main() {
	s := Switch("off")
	//var t Toggleable = s	// 报错，Toggleable接口里toggle()是*s的方法，对s应该取地址
	// GO 判断值是否满足一个接口的时候，指针方法并没有包含直接的值，但是它们包含指针。
	var t Toggleable = &s
	t.toggle()	// on
}
