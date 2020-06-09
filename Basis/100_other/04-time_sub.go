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

}
