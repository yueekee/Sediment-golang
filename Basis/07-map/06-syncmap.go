package main

import (
	"fmt"
	"sync"
)

/*
sync.Map 保证并发安全，但有一些性能损失
*/

func main() {
	var smap sync.Map

	// 存储该键值
	smap.Store(1, "one")
	smap.Store(2, "two")
	smap.Store(3, "three")
	smap.Store(4, "four")
	smap.Store(5, "five")

	// 获取该键对应的值
	fmt.Println(smap.Load(1))	// one true
	fmt.Println(smap.Load("1"))	// <nil> false

	// 删除该键对应的值
	smap.Delete(1)
	smap.Delete("1")

	// 遍历键值对
	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	/*
	3 three
	4 four
	5 five
	2 two
	*/

	// 加载并删除该元素
	fmt.Println(smap.LoadAndDelete(2))	// two true

	// 加载并存储该值
	fmt.Println(smap.LoadOrStore(6, "six"))

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	/*
	3 three
	4 four
	5 five
	6 six
	*/
}
