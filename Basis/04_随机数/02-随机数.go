package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 10; i++ {
	//	intn := rand.Intn(10)      // 生成0-9的随机数，不包含10
	//	fmt.Printf("%02d\n", intn) // 默认2为数字 01、02
	//}

	for i := 0; i < 30; i++ {
		s := fmt.Sprintf("%03d", rand.Intn(1000))
		fmt.Println(s)
	}
}


