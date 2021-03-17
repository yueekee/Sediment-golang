package main

import (
	"fmt"
	"sort"
)

/*总结：
map的遍历是无序的
但是直接打印map里是按顺序的
*/

func main() {
	ages := map[string]int{
		"a": 1,
		"c": 3,
		"d": 4,
		"b": 2,
	}

	fmt.Println(ages)
	fmt.Println(ages)
	fmt.Println(ages)
	//map[a:1 b:2 c:3 d:4]
	//map[a:1 b:2 c:3 d:4]
	//map[a:1 b:2 c:3 d:4]

	for k, v := range ages {
		fmt.Printf("%v: %v\n", k, v)
	}
	//b: 2
	//a: 1
	//c: 3
	//d: 4

	for k, v := range ages {
		fmt.Printf("%v: %v\n", k, v)
	}
	//a: 1
	//c: 3
	//d: 4
	//b: 2

	sortMap(ages)
	//1
	//2
	//3
	//4
}

func sortMap(ages map[string]int) {
	names := make([]string, 0, len(ages))

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names) //sort keys alphabetically

	for _, name := range names {
		fmt.Println(ages[name])
	}
}