package main

import "fmt"

/*总结
答案：
showA
showB
teacher showB
Teacher 没有自己 ShowA()，所以调用内部类型 People 的同名方法
*/

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA() // 只能使用People的方法
	t.ShowB() // 这个才是Teacher的方法
}

