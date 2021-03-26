package main

import "fmt"

/*总结：
大括号的作用域内，会出现变量隐藏的问题
*/

func main() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)	// 1
		i,x := 2,2
		fmt.Println(i,x) // 2 2
	}
	fmt.Println(x)  // 1

}
