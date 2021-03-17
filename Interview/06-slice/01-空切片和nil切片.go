package main

import (
	"fmt"
	"reflect"
)

/*总结：
slice只能和nil做比较，不能两个切片做比较
[]int 和 []int{} 是两种类型
*/

func main() {
	var s1 []int     // 为nil
	var s2 = []int{} // 为nil切片
	var n1 []int
	var	n2 = []int{1}
 	n1 = append(n1, 1)

	if s1 == nil {
		fmt.Println("s1 yes nil")
	}

	if s2 == nil {
		fmt.Println("s2 yes nil")
	}

	if reflect.DeepEqual(s1, s2) {	// 不相等
		fmt.Println("s1 equal s2")
	}

	if reflect.DeepEqual(n1, n2) { // 相等
		fmt.Println("n1 equal n2")
	}

	//if s1 == s2 {}	// 报错: slice can only be compared to nil
}
