package main

import (
	"fmt"
)

/* 总结：
	1.值.(数据类型)
	value, ok := v.(int)	// 返回值和该值是否为该数据类型的判断
	2.值.(type)
	switch value := v.(type)

*/
func main() {
	sli := make([]interface{}, 0)
	sli = append(sli, 12, "abc", 't', 12.34)
	for _, v := range sli {
		if value, ok := v.(int); ok {
			fmt.Println(value, "是int类型")
		} else if value, ok := v.(string); ok {
			fmt.Println(value, "是string类型")
		} else if value, ok := v.(rune); ok {
			fmt.Println(string(value), "是byte类型")
		} else if value, ok := v.(float64); ok {
			fmt.Println(value, "是float64类型")
		} else {
			fmt.Println(value, "未知类型")
		}
	}
	fmt.Println()
	// 第二种方法：
	for _, v := range sli {
		switch value := v.(type) {
		case int :
			fmt.Println(value, "是int类型")
		case string:
			fmt.Println(value, "是string类型")
		case rune:
			fmt.Println(string(value), "是byte类型")
		case float64:
			fmt.Println(value, "是float64类型")
		default:
			fmt.Println(value, "未知类型")
		}
	}
}