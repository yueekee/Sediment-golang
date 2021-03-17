package main

import "fmt"

/*总结：
切片的截取 [i:j]和[i:j:k]
如果i省略，默认为0；如果j省略，默认底层数组的长度；如果k省略，默认容量大小为底层数组中当前元素到数组末尾的长度
长度为j-i, 容量为k-i
*/

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	d := s[0:2]

	fmt.Printf("%v, len: %v, cap: %v\n", a, len(a), cap(a))
	fmt.Printf("%v, len: %v, cap: %v\n", b, len(b), cap(b))
	fmt.Printf("%v, len: %v, cap: %v\n", c, len(c), cap(c))
	fmt.Printf("%v, len: %v, cap: %v\n", d, len(d), cap(d))
	//[], len: 0, cap: 3
	//[1 2], len: 2, cap: 3
	//[2], len: 1, cap: 2
	//[1 2], len: 2, cap: 3
}