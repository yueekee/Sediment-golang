package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
		fmt.Println(v)
	}
}

/* 不会出现无限循环♻️
[1 2 3 0]
[1 2 3 0 1]
[1 2 3 0 1 2]
*/