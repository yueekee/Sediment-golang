package main

import "fmt"

/*总结：
MyInt1 为以int为基础类型的另一个类型，赋值时类型要对应
MyInt2 即int类型
*/

type MyInt1 int
type MyInt2 = int

func main() {
	var i int = 0
	//var i1 MyInt1 = i	// 报错：i不是int类型，而应该是MyInt1类型
	var y MyInt1 = 1
	var j MyInt1 = y
	var k MyInt1 = MyInt1(1)
	var l MyInt2 = i
	fmt.Println(j, k, l)
}
