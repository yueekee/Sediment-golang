package main

import (
	"fmt"
	"strings"
)

/*总结：
Count 返回字符串str中有几个不重复的的substr子串
*/

func main() {
	str := "abca"
	fmt.Println(strings.Count(str, ""))	// 5 如果substr是空，值为len（str）+1
	fmt.Println(strings.Count(str, "a"))	// 2
	fmt.Println(strings.Count(str, "ab"))	// 1
}
