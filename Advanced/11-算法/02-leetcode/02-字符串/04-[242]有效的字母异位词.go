package main

import (
	"fmt"
)

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
// 输入: s = "rat", t = "car"
//输出: false
//
// 说明:
//你可以假设字符串只包含小写字母。
//
// 进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// Related Topics 排序 哈希表


//leetcode submit region begin(Prohibit modification and deletion)
// 字母异位词，由颠倒字母顺序构成的词
// 说明每个字母出现的次数一样
func isAnagram1(s string, t string) bool {
	byteArr1 := make([]rune, 26)
	byteArr2 := make([]rune, 26)
	for _, v := range s {
		byteArr1[v - 'a']++
	}
	for _, v := range t {
		byteArr2[v - 'a']++
	}
	// 比较两个数组是否相同
	// 可以使用reflect进行比较，但是性能比较差
	37//equal := reflect.DeepEqual(byteArr1, byteArr2)
	equal := isSliceEqual(byteArr1, byteArr2)
	return equal
}

func isSliceEqual(a, b []rune) bool {
	// 比较长度是否相等
	if len(a) != len(b) {
		return false
	}

	// 判断两个slice是否都为nil或都不为nil
	if (a == nil) != (b == nil) {
		return false
	}

	// 比较对应索引处两个slice的值是否相等
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// 使用map处理，原理同上
func isAnagram(s string, t string) bool {
	bytemap1 := make(map[rune]int)
	bytemap2 := make(map[rune]int)
	for _, v := range s {
		bytemap1[v]++
	}
	for _, v := range t {
		bytemap2[v]++
	}
	//equal := reflect.DeepEqual(bytemap1, bytemap2)
	equal := isMapEqual(bytemap1, bytemap2)

	return equal
}

func isMapEqual(a, b map[rune]int) bool {
	// 比较长度是否相等
	if len(a) != len(b) {
		return false
	}

	// 比较两个map中key对应的值是否相同
	for k, av := range a {
		bv, ok := b[k]
		if !ok || bv != av {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)



func main() {
	anagram := isAnagram("anagram", "nagaram")
	fmt.Println(anagram)
}

