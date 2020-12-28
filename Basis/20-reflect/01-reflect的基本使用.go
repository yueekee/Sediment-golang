package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string "姓名"
	age  int    "年龄"
}

func main() {
	var stu = Student{"tom", 23}
	t := reflect.TypeOf(stu)  //反射对象的类型
	v := reflect.ValueOf(stu) //反射对象的值
	fmt.Println("t:", t)
	fmt.Println("v:", v)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("t.Field(i).Name:%v, v.Field(i):%v, t.Field(i).Tag:%v, t.Field(i).Type:%v\n",
			t.Field(i).Name, v.Field(i), t.Field(i).Tag, t.Field(i).Type)
	}
}
