package main

import "fmt"

/*总结：
切片的截取 [i:j]和[i:j:k]
如果i省略，默认为0；如果j省略，默认底层数组的长度；如果k省略，默认容量大小为底层数组中当前元素到数组末尾的长度
长度为j-i, 容量为k-i

copy(dst, src []Type) 将src切片复制到dst切片中
p.s.
使用char1[0:n]也可以实现这种方法(char1[0:n]左闭右开区间)
需要注意的是在遇到大切片扩容或者复制时可能会发生大规模的内存拷贝，一定要减少类似操作避免影响程序的性能
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

	// --------------
	char1 := []string{"a", "b", "c", "d", "e"}
	char2 := make([]string, 3)
	copy(char2, char1[1:])
	char3 := char1[1:3]

	fmt.Println("char1:", char1)    // char1: [a b c d e]
	fmt.Println("char2:", char2)    // char2: [b c ]	第三个元素为空
	fmt.Println("char3:", char3)    // char3: [b c]
	fmt.Printf("char1:%p\n", char1) // char1:0xc00019a000
	fmt.Printf("char2:%p\n", char2) // char2:0xc00018e030
	fmt.Printf("char3:%p\n", char3) // char3:0xc000072060
}
