package main

import (
	"fmt"
	"time"
)

// 时间转换
type Time time.Time

type Param struct {
	Start Time
}

func main() {
	var p Param
	add := time.Time(p.Start).Add(time.Hour * 24 * 365)
	fmt.Println(add)
}
