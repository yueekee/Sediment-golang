package main

import (
	"flag"
	"fmt"
	"time"
)

// 02_flag.Duration中第2个参数为value，可以通过-period命令行标志控制sleep时间
var period = flag.Duration("period", 1 * time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

//  ./03_flag_value -period 5s    // 睡眠5s 参数：1m，1h，50ms
