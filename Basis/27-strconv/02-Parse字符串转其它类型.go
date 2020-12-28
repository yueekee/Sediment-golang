package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

// Parse系列为将"字符串"转化为"其它类型"
func main() {
	// ParseFloat
	input := "12.5"
	grade, err := strconv.ParseFloat(input, 64)	// bitSize表示结果的精度位数
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(grade)	// 12.5

	// ParseBool
	// true: "1", "t", "T", "true", "TRUE", "True"
	// false: "0", "f", "F", "false", "FALSE", "False"
	input = "1"
	parseBool, err := strconv.ParseBool(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parseBool)	// true

	// ParseUint
	fmt.Println(math.MaxInt32)
	input = "21474836471"
	parseUint, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		log.Fatal(err)	// strconv.ParseUint: parsing "21474836471": value out of range
	}
	fmt.Println(parseUint)

	// ParseInt
	input = "123"
	parseInt, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parseInt)	// 123
}
