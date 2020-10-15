package main

import (
	"fmt"
	"time"
)

/*
给定一个整数数组，判断是否存在重复元素。

如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
*/

func main() {
	nums := []int{1, 2, 3, 4, 1}
	time1 := time.Now()
	containsDuplicate(nums)
	fmt.Println("time1 spent:", time.Since(time1))

}

func containsDuplicate1(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i] {
				return true
			}
		}
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > 1 {
			return true
		}
	}
	return false
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			return true
		} else {
			m[nums[i]] = 1
		}
	}
	return false
}
