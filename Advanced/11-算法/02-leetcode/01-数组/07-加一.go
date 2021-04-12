package main

import (
	"fmt"
)

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func main() {
	digits := []int{7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6}
	one := plusOne1(digits)
	fmt.Println(one)
}

func plusOne1(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func plusOne2(digits []int) []int {
	n := len(digits)
	index := 1
	for i := n - 1; i >= 0; i-- {
		sum := digits[i] + index //确定每一位数字的值
		value := sum % 10        //逢10归0,负责保持原数字
		index = sum / 10         //如果大于等于10,则为1;否则为0
		digits[i] = value
	}
	if index == 1 { //说明数组最前端是个10,但是1未放入数组
		digits = append([]int{1}, digits...)
	}
	return digits
}

// 转成字符串，有数组过长，超出int长度的情况，不好处理
//func plusOne(digits []int) []int {
//	var n int64
//	for _, v := range digits {
//		n = n*10 + int64(v)
//	}
//	fmt.Println("n:", n)
//	n += 1
//	s := strconv.FormatInt(n, 10)
//
//	ints := make([]int, len(s))
//	for i := len(s)-1; i >= 0; i-- {
//		ints[i] = int(n) % 10
//		n /= 10
//	}
//	return ints
//}
