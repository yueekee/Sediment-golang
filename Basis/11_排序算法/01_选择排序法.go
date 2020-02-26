package main

import "fmt"

/* 拿第n个数字和之后的所有数字进行比较，大的排到前面

 */

func main() {
	num := []int{20, 18, 25, 36, 15, 8, 12}
	for n := 0; n < len(num)-1; n++ {
		for i := n + 1; i <= len(num)-1; i++ {
			if num[i] > num[n] {
				num[i], num[n] = num[n], num[i]
			}
		}
	}
	fmt.Println(num)
}
