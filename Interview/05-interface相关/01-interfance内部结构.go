package main

import (
	"fmt"
)

/*总结：
当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil
live的动态值是nil，但是live的动态类型是*Student，是个nil指针，所以live不为nil
*/

var in interface{}

type People interface {
	Show()
}

type Student struct{}
func (stu *Student) Show() {}

func live() People {
	var stu *Student
	return stu
}

func main() {
	fmt.Println(live())        // nil
	fmt.Println(live() == nil) // false
	fmt.Println(in)            // nil
	fmt.Println(in == nil)     // true
}

