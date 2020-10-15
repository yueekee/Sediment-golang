package main

import (
	"fmt"
	"regexp"
	"strings"
)

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
// 示例 1:
//
// 输入: "A man, a plan, a canal: Panama"
//输出: true
//
//
// 示例 2:
//
// 输入: "race a car"
//输出: false
//
// Related Topics 双指针 字符串


//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome1(s string) bool {
	var r []rune
	for _, v := range s {
		// Ascii编码中0-9:[48-57], A-Z:[65-90], a-z:[97-122]
		if v < 48 || (v > 57 && v < 65) || v > 122 || (v > 90 && v < 97) {
			continue
		} else if v >= 97 && v <= 122 {
			r = append(r, v - 32)	// 转换为大写字母
		} else {
			r = append(r, v)
		}
	}
	fmt.Println("r:", r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	reg := regexp.MustCompile(`[a-zA-Z0-9]+`)
	allString := reg.FindAllString(s, -1)
	fmt.Println("allString:", allString)
	join := strings.Join(allString, "")
	fmt.Println("join:", join)
	r := strings.ToLower(join)
	fmt.Println("r:", r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	s := "A man, a plan, a canal: Panama"
	palindrome := isPalindrome(s)
	fmt.Println(palindrome)
}
