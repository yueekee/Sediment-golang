package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"d", "a", "ab", "c"}
	sort.Strings(s)
	fmt.Printf("%#v\n", s) // []string{"a", "ab", "c", "d"}

	n := []int{4, 2, 3}
	sort.Ints(n)
	fmt.Printf("%#v\n", n) // []int{2, 3, 4}
}
