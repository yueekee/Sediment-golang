package main

import (
	"fmt"
	"time"
)

func main() {
	// time.ParseInLocation 将string转化为当地的Time,并返回Time
	time1String := "2020-06-09 10:51:06"
	Time1, _ := time.ParseInLocation("2006-01-02 15:04:05", time1String, time.Local)
	Time2 := Time1.AddDate(0, 0, 1)
	DurationSub := Time2.Sub(Time1)

	// Truncate 返回最接近但早于Time1的时间点;如果DurationSub<0，则但会Time1的copy
	// Time1: 2020-06-09 10:51:06 +0800 CST
	DurationTruncate := Time1.Truncate(DurationSub)
	fmt.Println("DurationTruncate:", DurationTruncate) // DurationTruncate: 2020-06-09 08:00:00 +0800 CST 这里是按天算

	DurationTruncate2 := Time1.Truncate(time.Hour)
	fmt.Println("DurationTruncate2:", DurationTruncate2)	// DurationTruncate2: 2020-06-09 10:00:00 +0800 CST

	DurationTruncate3 := Time1.Truncate(time.Second)
	fmt.Println("DurationTruncate3:", DurationTruncate3)	// DurationTruncate3: 2020-06-09 10:51:06 +0800 CST

	DurationTruncate4 := Time1.Truncate(-time.Hour)
	fmt.Println("DurationTruncate4:", DurationTruncate4)	// DurationTruncate4: 2020-06-09 10:51:06 +0800 CST
}
