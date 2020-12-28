package main

import "fmt"

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

*/
func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}

// O(n^2)
func moveZeroes1(nums []int) {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		count++
		if nums[i] == 0 {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[len(nums)-1] = 0
			i--
		}
		if count > len(nums) {
			break
		}
	}
}

// O(n)
func moveZeroes2(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
		fmt.Println(i, ":", nums)
	}

	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[count] = nums[count], nums[i]
			count++
		}
	}
}