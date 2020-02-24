package main

import "fmt"

/* 总结：
	格式 value, ok := stu[1]
	value为stu[1]的值；ok为布尔类型，若存在则为true。

	delete()可以直接删除stu中的键值对。
*/

func main() {
	stu := map[int]string{1: "eric", 5: "loria", 3: "yoly"}

	if s, ok := stu[1]; ok {
		fmt.Println("值为", s)
	} else {
		fmt.Println("值不存在")
	}

	delete(stu, 5)
	fmt.Println(stu)
}
