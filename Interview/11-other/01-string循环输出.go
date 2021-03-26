package main

import "fmt"

/*总结：todo 没看懂
如果类型实现 String() 方法，当格式化输出时会自动使用 String() 方法。上面这段代码是在该类型的 String() 方法内使用格式化输出，导致递归调用，最后抛错。
*/

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	//fmt.Printf("%#v", c)
	return fmt.Sprintf("print: %v", c)
}

func main() {
	c := &ConfigOne{}
	c.String()
}

/*
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow
*/