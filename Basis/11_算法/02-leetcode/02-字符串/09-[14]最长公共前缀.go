package main

import (
	"fmt"
	"strings"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1:
//
// 输入: ["flower","flow","flight"]
//输出: "fl"
//
//
// 示例 2:
//
// 输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//
// 说明:
//
// 所有输入只包含小写字母 a-z 。
// Related Topics 字符串


//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	firstStr := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], firstStr) != 0 {
			firstStr = firstStr[:len(firstStr) - 1]
		}
	}
	if len(firstStr) > 0 {
		return firstStr
	}
	return ""
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	s := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(s))
}
