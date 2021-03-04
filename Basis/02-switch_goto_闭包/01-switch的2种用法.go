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
	var money int
	switch1(money)
	//switch2(money)
}

func switch1(money int) {
	for {
		fmt.Print("输入金额:")
		fmt.Scan(&money)	// 获取键盘输入的值
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
}

func switch2(money int) {
	for {
		fmt.Println("输入金额:")
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
	}
}
