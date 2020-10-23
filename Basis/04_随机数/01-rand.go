package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/* 总结：
使用时间戳作为随机种子，生成随机数。
若种子一样，生成的为伪随机数
*/

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
		fmt.Println(rand.Intn(100))
	}

	unix := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println("unix:", unix)
	p := unix + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10))
	fmt.Println("p:", p)
}
