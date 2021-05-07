package main

import "fmt"

/*总结：
string转[]byte是值拷贝
*/

func main() {
	s := "a"
	fmt.Printf("%p\n", &s)
	b := []byte(s)
	fmt.Printf("%p\n", b)
}
