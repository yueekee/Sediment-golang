package main

import "fmt"

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c.String())
}

func main() {
	c := &ConfigOne{}
	c.String()
}
