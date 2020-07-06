package main

import (
	"fmt"
	"os"
	"strconv"
)

//给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
//
// 注意：整数序列中的每一项将表示为一个字符串。
//
// 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
//
// 1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//
//
// 第一项是数字 1
//
// 描述前一项，这个数是 1 即 “一个 1 ”，记作 11
//
// 描述前一项，这个数是 11 即 “两个 1 ” ，记作 21
//
// 描述前一项，这个数是 21 即 “一个 2 一个 1 ” ，记作 1211
//
// 描述前一项，这个数是 1211 即 “一个 1 一个 2 两个 1 ” ，记作 111221
//
//
//
// 示例 1:
//
// 输入: 1
//输出: "1"
//解释：这是一个基本样例。
//
// 示例 2:
//
// 输入: 4
//输出: "1211"
//解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 = 2；类似
//"1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。
// Related Topics 字符串


//leetcode submit region begin(Prohibit modification and deletion)
func countAndSay(n int) string {
	// 计算连着的数字有几个
	// 采用递归
	if n == 1 {
		return "1"
	} else if n == 2 {
		return "11"
	}
	s := countAndSay(n - 1)
	fmt.Printf("++++n:%v, s:%v\n", n, s)
	count := 1
	r := ""
	/*
	这里为什么要从i := 1开始？
	是因为i++是在一次循环完成之后才进行执行的，如果i从0开始，if i == len(s) - 1的判断就有问题
	*/
	for i := 1; i <= len(s)-1; i++ {
		if s[i-1] != s[i] {
			r += strconv.Itoa(count) + string(s[i-1])
			count = 1
			fmt.Println("----s[i] != s[i+1]的r：", r)
		} else {
			count ++
		}
		if i == len(s) - 1 {
			r += strconv.Itoa(count) + string(s[i])
			fmt.Println("----队尾r：", r)
			s = r
			return s
		}
	}
	return s
}
//leetcode submit region end(Proh		ibit modification and deletion)

/*
	n=1, 1
	n=2, 11
	n=3, 21X
	n=4, 1211
	n=5, 111221
*/
func main() {
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Println(countAndSay(n))
}
