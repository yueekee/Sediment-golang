package main

import (
	"fmt"
	"reflect"
)

type data struct {
	num    int
	checks [10]func() bool
	doit   func() bool
	m      map[string]string
	bytes  []byte
}

func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2))
	// v1 == v2: true

	m1 := map[int]string{1:"1", 2:"2"}
	m2 := map[int]string{2:"2", 1:"1"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))
	// m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	// s1 == s2: true
}
