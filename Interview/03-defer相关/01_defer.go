package main

import "fmt"

/*	注意：
1.defer执行的是最外层函数，内层的函数还是会被先调用的；
2.defer执行的函数的变量已经确定，不再改变。
*/

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
