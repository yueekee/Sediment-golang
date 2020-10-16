package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().String())
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Local())

	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	rand.Seed(time.Now().UnixNano())
	timeStamp = timeStamp + strconv.Itoa(rand.Intn(1000))

}
