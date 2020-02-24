package main

import "fmt"

/* 总结：
	1.goto跳转到标签
	2.标签需要使用:
*/

func main() {
	i := 1
	label1:
	fmt.Println(i)
	i++

	for i <= 5 {
		goto label1
	}
	//breakLabel()
}

// 另外中断多重循环也可使用“标签”
func breakLabel() {
	label:
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if i == 5 {
				break label		// 中断到标签处
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}