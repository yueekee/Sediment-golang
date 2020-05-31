package main

import "fmt"

/* copy函数
p.s.
使用char1[0:n]也可以实现这种方法
char1[0:n]左闭右开区间
*/

func main() {
	char1 := []string{"a", "b", "c", "d", "e"}
	char2 := make([]string, 3)
	//char2 = char1[:3]
	copy(char2, char1[1:3])

	fmt.Println("char1:", char1)
	fmt.Printf("char1:%p\n", char1)
	fmt.Println("char2:", char2)
	fmt.Printf("char2:%p\n", char2) // 不同地址
}
