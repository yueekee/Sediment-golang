package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println(time.Now())                               // 2020-12-15 16:11:18.061577 +0800 CST m=+0.000102257
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 2020-12-15 16:11:18
	fmt.Println(time.Now().Second())                      // 18
	fmt.Println(time.Now().String())                      // 2020-12-15 16:11:18.061745 +0800 CST m=+0.000269899
	fmt.Println(time.Now().Nanosecond())                  // 61748000
	fmt.Println(time.Now().Unix())                        // 1608019878
	fmt.Println(time.Now().UnixNano())                    // 1608019878061751000
	fmt.Println(time.Now().Local())                       // 2020-12-15 16:11:18.061752 +0800 CST

	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	rand.Seed(time.Now().UnixNano())
	timeStamp = timeStamp + strconv.Itoa(rand.Intn(1000))

}
