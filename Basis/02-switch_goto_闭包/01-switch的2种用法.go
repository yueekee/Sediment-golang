package main

import "fmt"

/* 总结：
switch有2中用法：
	1.switch {
		case 条件1：
			条件语句1
		case 条件2：
			条件语句2
		default：
			默认条件语句
		}
	2.switch 条件语句{
		case 条件语句对应的值1：
			代码1
		case 条件语句对应的值2：
			代码2
		default：
			默认代码
		}
*/
func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("---------")
	switch {
	case i == 1:
		fmt.Println("one")
	case i == 2:
		fmt.Println("two")
	case i == 3:
		fmt.Println("three")
	}
}

