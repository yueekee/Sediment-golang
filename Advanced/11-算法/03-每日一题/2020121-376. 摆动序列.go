package main

func main() {
	
}

/*
输入: [1,17,5,10,13,15,10,5,16,8]
输出: 7
解释: 这个序列包含几个长度为 7 摆动序列，其中一个可为[1,17,10,13,10,16,8]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/wiggle-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
var nums = []int{1,17,5,10,13,15,10,5,16,8}

// 1,17,5,10,13,15,10,5,16,8
// 1    5    13    10   16
//   17   10    15    5    8

func wiggleMaxLength(nums []int) int {
	//var flag int = 1
	if nums[0] - nums[1] < 0 {
		//flag = -1
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] - nums[j] > 0 {

			}
		}
	}
}