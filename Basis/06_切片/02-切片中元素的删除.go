package main

import "fmt"

func main() {
	sli := []int{0,0,0,1,2,3,4,5,6,7,8,9}
	n := 0
	ints := RemoveParam(sli, n)
	fmt.Println(ints)
}

func RemoveParam(sli []int, n int) []int {
	for i := 0; i < len(sli); i++ {
		if sli[i] == n {
			if i == 0 {
				sli = sli[1:]
				i-- // 如果索引0被去掉后，原来索引1的会移动到索引0
			} else if i == len(sli) - 1 {
				sli = sli[:i]
			} else {
				sli = append(sli[:i], sli[i+1:]...)
			}
		}
	}
	return sli
}
