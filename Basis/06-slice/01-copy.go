package main

import "fmt"

/* copy(dst, src []Type) 将src切片复制到dst切片中
p.s.
使用char1[0:n]也可以实现这种方法(char1[0:n]左闭右开区间)
需要注意的是在遇到大切片扩容或者复制时可能会发生大规模的内存拷贝，一定要减少类似操作避免影响程序的性能
*/

func main() {
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
