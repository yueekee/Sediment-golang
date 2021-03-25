package main

import "fmt"

/*总结：
答案为5
hello()函数的参数在执行defer语句的时候会保存一份副本，在实际调用hello()函数时使用
*/

func hello(i int) {
	fmt.Println(i)
}

func main() {
	i := 5
	defer hello(i)
	i += 10
}
