package main

import "fmt"

/*总结：make和new的区别
make 初始化内置的类型，只适用于slice，map，channel
new 创建指定类型的指针空间
*/

func main() {
	m := make(map[int]string)
	//m[1] = "1"
	fmt.Printf("%#v\n", m)	// map[int]string{}

	var n *map[int]string
	n = new(map[int]string)
	fmt.Printf("%#v\n", n)	// &map[int]string(nil)

	s := make([]int, 0)
	fmt.Printf("%#v\n", s)	// []int{}

	var s2 *[]int
	s2 = new([]int)
	fmt.Printf("%#v\n", s2)	// &[]int(nil)
}
