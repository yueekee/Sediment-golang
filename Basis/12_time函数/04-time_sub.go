package main

import (
	"fmt"
	"time"
)

func main() {

	payDate := "2020-06-09 10:51:06"
	t := time.Now()
	location, _ := time.ParseInLocation("2006-01-02 15:04:05", payDate, time.Local)
	t2 := location.AddDate(0, 0, 1)
	sub := t2.Sub(t)
	fmt.Println(sub.Seconds())
	fmt.Println("___________")


	startTime := "1970-01-01 00:00:00"
	parse, _ := time.Parse("2006-01-02 15:04:05", startTime) // 格式要求比较高
	fmt.Println("+++parse:", parse)
	endTime := "2100-06-05 00:00:00"
	parse2, _ := time.Parse("2006-01-02 15:04:05", endTime)
	fmt.Println("+++parse2:", parse2)
	if parse.After(parse2) {
		fmt.Println("111")
	}
}
