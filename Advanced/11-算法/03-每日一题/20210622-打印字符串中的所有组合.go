package main

import "fmt"

/*
输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]
*/

func main() {
	s := "abc"
	strings := permutation(s)
	fmt.Println(strings)
}

func permutation(s string) []string {
	ss := make([]string, 1)
	if len(s) == 0 {
		return nil
	}

	if len(s) == 1 {
		ss[0] = s
		return ss
	}

	// ab-> 2种
	// abc->3*2种
	// aba->3种
	// abcd->4*3*2=24种
	// todo 这种算法是什么

	return nil
}
