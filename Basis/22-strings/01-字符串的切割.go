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
}
