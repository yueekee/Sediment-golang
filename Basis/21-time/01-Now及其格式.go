package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now())                               // 2020-12-15 16:11:18.061577 +0800 CST m=+0.000102257
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 2020-12-15 16:11:18
	fmt.Println(time.Now().Second())                      // 18
	fmt.Println(time.Now().String())                      // 2020-12-15 16:11:18.061745 +0800 CST m=+0.000269899
	fmt.Println(time.Now().Nanosecond())                  // 61748000	共8位数
	fmt.Println(time.Now().Unix())                        // 1608019878	共10位数
	fmt.Println(time.Now().UnixNano())                    // 1608019878061751000 共19位数
	fmt.Println(time.Now().Local())                       // 2020-12-15 16:11:18.061752 +0800 CST

	rand.Seed(time.Now().UnixNano()) // UnixNano()常被用来做随机数种子
	fmt.Println(rand.Intn(1000))     // [0,1000)的随机数
}
