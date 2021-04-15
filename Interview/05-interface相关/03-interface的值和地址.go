package main

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	//g(s) //B
	f(p) //C
	//g(p) //D
}

/*
golang是强类型语言（变量定义好类型，类型就不能再改变），interface是所有类型的父类，
f(x interface{})支持传入所有类型
而g(x *interface{})只能接收*interface{}
*/
