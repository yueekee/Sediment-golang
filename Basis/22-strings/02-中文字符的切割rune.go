package main

import "fmt"

func main() {
	info := "史蒂夫史蒂夫1"
	infoRune := []rune(info)
	s := string(infoRune[:4])
	fmt.Println(s)
}
