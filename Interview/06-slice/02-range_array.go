package main

import "fmt"

/*总结：
range 表达式是副本参与循环，并不是原来的数组。针对数组有这个问题，对slice却没有这个问题
因为slice有一个指向底层数组的指针，当range表达式发生复制时，副本的指针依旧指向原底层数组
*/

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	/*
	r =  [1 2 3 4 5]
	a =  [1 12 13 4 5]
	*/
	fix()
	/*
	r =  [1 12 13 4 5]
	a =  [1 12 13 4 5]
	*/
	fmt.Println("---------")
	fix2()
	/*
	r =  [1 12 13 4 5]
	a =  [1 12 13 4 5]
	*/
	testSlice()
	/*
	r =  [1 12 13 4 5]
	a =  [1 12 13 4 5]
	*/
}

func fix() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

func fix2() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i:= 0; i < len(a) ; i++ {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = a[i]
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}


func testSlice() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	/*
	r =  [1 12 13 4 5]
	a =  [1 12 13 4 5]
	*/
}