package main

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{4,1,2,1,2}
	number := singleNumber(nums)
	fmt.Println(number)
}

func singleNumber1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] ++
	}
	for i, i2 := range m {
		if i2 == 1 {
			return i
		}
	}
	return 0
}

func singleNumber(nums []int) int {
	n := 0
	for _, num := range nums {
		n ^= num
	}
	return n
}
