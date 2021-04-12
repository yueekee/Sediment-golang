package main

import "fmt"

/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{0, 1, 1, 1, 2, 2, 3, 3, 4}
	duplicates := removeDuplicates1(nums)
	fmt.Println(duplicates)	// 5
}

// O(n^2)
func removeDuplicates1(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if j == len(nums)-1 {
					nums = nums[:j]
				} else {
					nums = append(nums[:j], nums[j+1:]...)
				}
				j--
			}
		}
	}
	fmt.Println(nums)	// [0 1 2 3 4]
	return len(nums)
}