package main

import "fmt"

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c.String())
}

func main() {
	//c := &ConfigOne{}
	//c.String()

	status := 1	// 0001
	num := 2 // 0010  3-> 0011
	if true {
		num |= status
	}
	fmt.Println("num:", num)
}
