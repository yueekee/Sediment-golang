package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一 个 goroutine 打印字母， 最终效果如下:
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

func main() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		fmt.Println("--go1--")
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()

	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		fmt.Println("--go2--")
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "") {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wait)
	number <- true
	fmt.Println("---1---")
	wait.Wait()
}

/*
这里用到了两个 channel 负责通知，letter负责通知打印字母的goroutine来打印字母， number用来通知打印数字的goroutine打印数字。
wait 用来等待字母打印完成后退出循环。
*/