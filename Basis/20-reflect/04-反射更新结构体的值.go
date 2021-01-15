package main

import (
	"reflect"
)

type A struct {
	Field1 string
	*B
}

type B struct {
	Field2 string
}

func main() {
	a := A{B: &B{Field2: "123"}}
	println(a.B.Field2) // 123
	v := reflect.ValueOf(&a)
	v.Elem().FieldByName("B").Elem().FieldByName("Field2").Set(reflect.ValueOf("456"))
	println(a.B.Field2) // 456
}
