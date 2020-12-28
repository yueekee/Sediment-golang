package main

import (
	"fmt"
	"reflect"
)

func main() {
	c := C{}
	f := reflect.ValueOf(c).MethodByName("Hello")
	f.Call(nil)
}

type C struct{}

func (c C) Hello() {
	fmt.Println("hello")
}
