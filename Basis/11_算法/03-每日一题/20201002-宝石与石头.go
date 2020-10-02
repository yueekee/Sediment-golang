package main

import "fmt"

/*
 给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。 S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jewels-and-stones
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

注意:

S 和 J 最多含有50个字母。
 J 中的字符不重复。
*/

func main() {
	J := "aAb"
	S := "aAAabbbb"
	stonesNum := numJewelsInStones2(J, S)
	fmt.Println(stonesNum)
}

func numJewelsInStones(J string, S string) int {
	m := make(map[int32]int)
	for _, i := range J {
		m[i] = 0
	}
	count := 0
	for _, i := range S {
		if _, ok := m[i]; ok {
			count ++
		}
	}

	return count
}

func numJewelsInStones2(J string, S string) int {
	count := 0
	for _, i2 := range J {
		for _, i4 := range S {
			if i2 == i4 {
				count ++
			}
		}
	}
	return count
}
