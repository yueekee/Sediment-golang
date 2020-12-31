package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Format系列为将"字符串"转化为"其它类型"
func main() {
	// FormatUint uint64 -> string
	var ID uint64 = 10
	idStr := strconv.FormatUint(ID, 10) // uint64 -> string, base表示进制
	fmt.Println(idStr)
	fmt.Println(reflect.TypeOf(idStr)) // string
	// 另一种方法
	idStr = strconv.Itoa(int(ID))
	fmt.Println(reflect.TypeOf(idStr)) // string

	// FormatInt int64 -> string
	var num int64 = 10
	formatInt := strconv.FormatInt(num, 10)
	fmt.Println(formatInt)

	// FormatFloat float64 -> string
	var pi float64 = 3.1415
	float := strconv.FormatFloat(pi, 'b', 1, 64)	// todo 参数的意义
	fmt.Println(float)

	var b bool = false
	formatBool := strconv.FormatBool(b)
	fmt.Println(formatBool)
}
