package main

import (
	"fmt"
	"time"
)

/* 总结：
定时器：使用time.After函数，多长时间后执行命令
时间周期：使用time.NewTicker
*/

func main() {
	timer := time.NewTimer(time.Second) // 1s后将时间写入通道
	fmt.Println(<-timer.C)              // 2020-02-25 16:31:49.2972707 +0800 CST m=+0.006019401

	timer2 := time.After(time.Second) // 和上面功能相同，使用After更为直观方便
	fmt.Println(<-timer2)

	ticker := time.NewTicker(time.Second) // 每隔1s进行一次
	for {
		<-ticker.C
		fmt.Println("这是ticker的打印")
	}
}
