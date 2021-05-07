package main

import (
	"fmt"
	"sync"
)

/* WaitGroup
1.新建WaitGroup组
2.Add添加任务数
3.每执行完一个任务，使用Done减去一个任务数
4.使用wait等待任务数为0时结束
*/

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(9)
	for i := 1; i < 10; i++ {
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(&waitGroup, i)
	}
	waitGroup.Wait()
}

// 输出为乱序的1-9位数
