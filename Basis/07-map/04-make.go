package main

import "fmt"

/*总结：make和new的区别
make 初始化内置的类型，只适用于slice，map，channel
new 创建指定类型的指针空间
*/

func main() {
	m := make(map[int]string)
	m[1] = "1"
	fmt.Println(m)	// map[1:1]

	var n *map[int]string
	n = new(map[int]string)
	fmt.Println(n)	// &map[]
}
