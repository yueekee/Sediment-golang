package main

import "fmt"

/*总结：一级函数
h 为hello的一级函数
y 为函数hello本身，其值为hello()的返回值
*/

func hello() []string {
	return nil
}

func main() {
	h := hello
	fmt.Printf("h: %#v\n", h)	// h: (func() []string)(0x10a89c0)
	y := hello()
	fmt.Printf("y: %#v\n", y)	// y: []string(nil)
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}