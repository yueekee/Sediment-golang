package main

import "fmt"

func main() {
	s := Square{len: 5}
	fmt.Println(s.Sides()) // 4
}

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 4
}

func (s *Square) Area() int {
	return s.len * s.len
}

// 定义为nil的Square指针类型的变量使用_进行接收
// 注意: 如果*Square没有实现完Shape接口内的所有方法，下面的声明就会出现问题
var _ Shape = (*Square)(nil)
