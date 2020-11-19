package _4_Map相关

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/* 注意：
1.main函数中map需要初始话才能用（panic: assignment to entry the_go_pro_lan nil map）
2.读写map需要进行加锁，否则可能会线程冲突(fatal error: concurrent map read and map write)

*/

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	//ua.Lock()
	//defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	u := &UserAges{}
	//u.ages = make(map[string]int)
	for i := 1; i <= 10; i++ {
		go func(i int) {
			u.Add("eric"+strconv.Itoa(i), 27)
		}(i)
	}
	for i := 1; i <= 100; i++ {
		go func(i int) {
			u.Get("eric")
		}(i)
	}

}

var lock sync.Mutex

func mapSafe() {
	m := make(map[int]int)
	go func() { //开一个协程写map
		for i := 0; i < 10000; i++ {
			lock.Lock() //加锁
			m[i] = i
			lock.Unlock() //解锁
		}
	}()
	go func() { //开一个协程读map
		for i := 0; i < 10000; i++ {
			lock.Lock() //加锁
			fmt.Println(m[i])
			lock.Unlock() //解锁
		}
	}()
	time.Sleep(time.Second * 20)

	lock.Lock()
	for i := 0; i < 10000; i++ {
		fmt.Println(m[i])
	}
	lock.Unlock()
}
