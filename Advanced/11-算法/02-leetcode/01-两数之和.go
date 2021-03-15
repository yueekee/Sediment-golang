//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。 
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。 
//
// 
//
// 示例: 
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
// 
// Related Topics 数组 哈希表
package main

import "fmt"
//leetcode submit region begin(Prohibit modification and deletion)
func twoSum1(nums []int, target int) []int {
	var arr []int
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums) - j; j++ {
			if nums[i] + nums[j] == target {
				arr = append(arr, i, j)
			}
		}
	}
	return arr
}

func twoSum2(nums []int, target int) []int {
	var arr []int
	hmap := make(map[int]int)	// map[2]0, map[7]1
	for i := 0; i < len(nums); i++ {
		hmap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := hmap[target - nums[i]]; ok {
			arr = append(arr, i, hmap[target - nums[i]])
			return arr
		}
	}
	return arr
}

// todo 是否可以在写map时，同时进行查询
func twoSum(nums []int, target int) []int {
	//var arr []int
	hmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hmap[nums[i]] = target - nums[i]
		if nums[i] + hmap[nums[i]] == target {
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	sum := twoSum(nums, target)
	fmt.Println("sum:", sum)
}