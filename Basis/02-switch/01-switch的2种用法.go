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

这里有个问题：输入金额为-09时，金额为0，感觉是取值只能拿到一位数。
*/
func main() {
	var money int
	for {
		fmt.Println("输入金额")
		fmt.Scan(&money)
		fmt.Println("输入金额：", money)
		switch {
		case money >= 4000:
			fmt.Println("钻石用户")
		case money >= 3000:
			fmt.Println("白金用户")
		case money >= 2000:
			fmt.Println("黄金用户")
		case money >= 1000:
			fmt.Println("白银用户")
		case money > 0:
			fmt.Println("普通用户")
		default:
			fmt.Println("请输入正确的金额")
		}
	}

	/*	for {
		fmt.Println("输入金额")
		fmt.Scan(&money)
		if money >= 4000 {
			fmt.Println("钻石用户")
		} else if money >= 3000 {
			fmt.Println("白金用户")
		} else if money >= 2000 {
			fmt.Println("黄金用户")
		} else if money >= 1000 {
			fmt.Println("白银用户")
		} else {
			fmt.Println("普通用户")
		}
	}*/
}
