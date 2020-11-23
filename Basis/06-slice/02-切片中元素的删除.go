package main

import "fmt"

func main() {
	sli := []int{0,0,0,1,2,3,4,5,6,7,8,9}
	n := 8
	ints := RemoveParam(sli, n)
	fmt.Println(ints)
}

func RemoveParam(sli []int, n int) []int {
	for i := 0; i < len(sli); i++ {
		if sli[i] == n {
			if i == 0 {
				sli = sli[1:]
			} else if i == len(sli) - 1 {
				sli = sli[:i]
			} else {
				sli = append(sli[:i], sli[i+1:]...)
			}
			i-- // 如果索引i被去掉后，原来索引i+1的会移动到索引i
		}
	}
	return sli
}
