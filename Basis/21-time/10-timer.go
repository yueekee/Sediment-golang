package main

import (
	"fmt"
	"time"
)

/* 总结：
定时器：使用time.NewTimer或time.After函数，多长时间后执行命令
时间周期：使用time.NewTicker
*/

func main() {
	timer := time.NewTimer(time.Second) // 1s后将时间写入通道
	fmt.Println(<-timer.C)              // 2020-02-25 16:31:49.2972707 +0800 CST m=+0.006019401

	timer2 := time.After(time.Second) // 和上面功能相同，使用After更为直观方便
	fmt.Println(<-timer2)

	ticker := time.NewTicker(time.Second) // 每隔1s进行一次
	i := 0
	for i < 5 {
		<-ticker.C
		fmt.Printf("这是ticker的打印:%d\n", i)
		i ++
	}
}

/*
2021-04-08 10:02:43.074451 +0800 CST m=+1.004562085
2021-04-08 10:02:44.076154 +0800 CST m=+2.006251184
这是ticker的打印:0
这是ticker的打印:1
这是ticker的打印:2
这是ticker的打印:3
这是ticker的打印:4
*/
