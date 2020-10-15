package main

import "fmt"

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。
*/
func main() {
	nums := []int{1,2,3,4,5,6,7}
	k := 3
	rotate(nums, k)
	fmt.Println("end: ", nums)
}

func rotate(nums []int, k int) {
	l := len(nums)
	if k <= 0 {
		return
	}
	tmp := 0
	for i := 1; i <= k; i++ {
		tmp = nums[l-1]
		for i := l-1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = tmp
		fmt.Println("i: ", nums)
	}
}
