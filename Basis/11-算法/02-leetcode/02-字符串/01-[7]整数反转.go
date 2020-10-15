package main

import (
	"fmt"
	"math"
	"strconv"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
// 示例 1:
//
// 输入: 123
//输出: 321
//
//
// 示例 2:
//
// 输入: -123
//输出: -321
//
//
// 示例 3:
//
// 输入: 120
//输出: 21
//
//
// 注意:
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231, 231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
// Related Topics 数学


//leetcode submit region begin(Prohibit modification and deletion)
func reverse1(x int) int {
	sign := 1
	y := 0
	if x < 0 {
		sign = -1
		x = x * sign
	}

	for x > 0 {
		y = y * 10 + x % 10
		x /= 10
	}
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	return y * sign
}

// 转换成字符串，反转字符，然后在转换为数字。比较挫的方法
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = x * sign
	}
	str := strconv.Itoa(x)
	s := []rune(str)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	i, _ := strconv.Atoi(string(s))
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	return i
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	x := 123
	i := reverse(x)
	fmt.Println("i:", i)
}
