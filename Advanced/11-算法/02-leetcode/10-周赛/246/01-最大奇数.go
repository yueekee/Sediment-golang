package main

import (
	"fmt"
	"strconv"
)

func main() {
	//num := "52"
	//num := "10133890"
	num := "23953767242388496965328710"
	number := largestOddNumber(num)
	fmt.Println("----:", number)

}

func largestOddNumber(num string) string {
	for i := len(num)-1; i >= 0 ; i-- {
		atoi, err := strconv.Atoi(string(num[i]))
		if err != nil {
			return ""
		}
		fmt.Println("atoi:", atoi)

		if atoi%2 != 0 {
			return num[:i+1]
		}
	}
	return ""
}

