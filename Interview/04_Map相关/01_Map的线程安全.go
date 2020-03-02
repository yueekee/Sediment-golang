package _4_Map相关

import (
	"strconv"
	"sync"
)

/* 注意：
1.main函数中map需要初始话才能用（panic: assignment to entry in nil map）
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
