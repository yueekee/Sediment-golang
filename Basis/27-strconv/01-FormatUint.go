package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var ID uint64 = 10

	id := strconv.FormatUint(ID, 10)	// uint64 -> string, base表示进制
	fmt.Println(id)
	fmt.Println(reflect.TypeOf(id)) // string

	id = strconv.Itoa(int(ID))
	fmt.Println(reflect.TypeOf(id)) // string
}
