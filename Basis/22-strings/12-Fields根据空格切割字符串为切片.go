package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

	// Fields are: ["foo" "bar" "baz"]
	// %q     单引号围绕的字符字面值，由Go语法安全地转义
}
