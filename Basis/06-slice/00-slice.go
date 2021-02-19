package main

import "fmt"

/*初始化
1。通过下标的方式获得数组或者切片的一部分；
2。使用字面量初始化新的切片；
3。使用关键字 make 创建切片：
*/
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &array) // 0xc0000160c0
	s := array[1:4]
	fmt.Printf("%#v, %p, len:%v, cap:%v\n", s, s, len(s), cap(s)) // []int{2, 3, 4}, 0xc0000160c8, len:3, cap:4

	slice := []int{1, 2, 3}
	fmt.Printf("%#v, len:%v, cap:%v\n", slice, len(slice), cap(slice)) // []int{1, 2, 3}, len:3, cap:3

	slice2 := make([]int, 3)
	fmt.Printf("%#v, len:%v, cap:%v\n", slice2, len(slice2), cap(slice2)) // []int{0, 0, 0}, len:3, cap:3

	// 一般声明长度为0的slice
	slice = make([]int, 0)
	// 在使用append进行添加元素
	slice = append(slice, 1, 2, 3)

}
