package main

import (
	"fmt"
	"reflect"
)

type worker struct {
	Func interface{}
	Args []reflect.Value
}

func main() {
	wk := worker{
		Func: func(x, y int) {
			fmt.Println(x,"+",y,"=",x + y)
		},
		Args: []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)},
	}

	reflect.ValueOf(wk.Func).Call(wk.Args)
}
