package main

import "fmt"

// 延迟对recover()的调用，不会panic，会正常结束
// panic后的代码无法再被执行

func calmDown() {
	recover()
}

func freakOut() {
	defer calmDown()
	panic("oh no")
	fmt.Println("I won't be run")	// panic后的这段代码永远不会被执行
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
