package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "asdfsa|sdf|sdf"
	trim := strings.Split(s, "|")
	fmt.Println("len:", len(trim))
	fmt.Println("trim", trim)
	fmt.Println("trim[0]", trim[0])


	s2 := "2020102114505434/20201021141122138_1.txt"
	fmt.Println("s2:", s2[36:])

	s3 := "2020102114505434/1.txt"
	if len(s3) >= 36 {
		fmt.Println("s3:", s3[36:])
	}


}
