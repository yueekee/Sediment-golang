package main

import (
	"fmt"
	"strings"
)

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//
//
// 示例：
//
// s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2
//
//
//
//
// 提示：你可以假定该字符串只包含小写字母。
// Related Topics 哈希表 字符串


//leetcode submit region begin(Prohibit modification and deletion)
// go中有strings.Index函数可以直接使用
func firstUniqChar1(s string) int {
	for _, v := range s {
		fIndex := strings.Index(s, string(v))
		lIndex := strings.LastIndex(s, string(v))
		if fIndex == lIndex {
			return fIndex
		}
	}
	return -1
}

// 使用数组，利用字母在ascii码中对应的值作为index
func firstUniqChar2(s string) int {
	b := make([]rune, 126)
	for _, v := range s {
		b[v]++
	}
	for k, v := range s {
		if b[v] == 1 {
			return k
		}
	}
	return -1
}

// 因为指定是小写字母，使用数组更节省空间
func firstUniqChar2_1(s string) int {
	b := [26]rune{}
	for _, v := range s {
		b[v - 'a']++
	}
	for k, v := range s {
		if b[v - 'a'] == 1 {
			return k
		}
	}
	return -1
}

// 使用map，原理同上
func firstUniqChar3(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for k, v := range s {
		if m[v] == 1 {
			return k
		}
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	s := "lovveleetcode"
	f := firstUniqChar(s)
	fmt.Println(f)
}
