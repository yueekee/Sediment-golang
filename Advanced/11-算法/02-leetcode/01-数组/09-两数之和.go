package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2jrse/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{3,2,4}
	target := 6
	ints := twoSum(nums, target)
	fmt.Println(ints)
}

// O(n^2)
func twoSum1(nums []int, target int) []int {
	ints := make([]int, 0)
	flag := false
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			//fmt.Printf("i:%d, j:%d\n", i, j)
			if nums[i] + nums[j] == target {
				ints = append(ints, i, j)
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	return ints
}

// O(2n)
func twoSum(nums []int, target int) []int {
	ints := make([]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	fmt.Println(m)

	for i := 0; i < len(nums); i++ {
		fmt.Println("i:", i)
		_, ok := m[target-nums[i]]
		if ok && m[target-nums[i]] != i {	// 需要判断target-nums[i]后的值和nums[i]相等的情况
			ints = append(ints, i, m[target-nums[i]])
			break
		}
	}

	return ints
}
