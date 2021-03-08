package main

import (
	"fmt"
)

/*总结：复制切片/数组元素 -> 指针切片出现的问题
遍历切片a元素，将其地址存到切片指针b，切片指针b中所有的元素指向切片a的地址，对应的值为切片最后的一个元素，见test1函数
解决办法：
1.将遍历出来的切片a的元素重新赋值——产生了新的地址，见test2函数
2.使用channel进行传递也可以不用担心值不对的问题，见goroutineTest函数
*/

func main() {
	test1()
	test12() // 也可以解决
	test2()
	goroutineTest()
}

func test1() {
	in := []int{1, 2, 3}

	var out []*int
	for i, _ := range in {
		out = append(out, &in[i])
	}

	fmt.Println("*out:", *out[0], *out[1], *out[2]) // *out: 3 3 3
	fmt.Println("out:", out[0], out[1], out[2]) // out: 0xc00001a0a8 0xc00001a0a8 0xc00001a0a8
	// 原因：每次迭代中，都会把v的地址附加到out切片中，但v的地址都是一样的，该地址都指向同一个值
}

func test12() {
	in := []int{1, 2, 3}

	var out []*int
	for i, _ := range in {
		out = append(out, &in[i])
	}

	fmt.Println("*out:", *out[0], *out[1], *out[2]) // *out: 1 2 3
	fmt.Println("out:", out[0], out[1], out[2])
}

func test2() {
	in := []int{1, 2, 3}

	var out []*int
	for _, v := range in {
		v := v	// 重新赋值，产生了新的地址，然后再附加到out切片中
		out = append(out, &v)
	}

	fmt.Println("*out:", *out[0], *out[1], *out[2]) // *out: 1 2 3
	fmt.Println("out:", out[0], out[1], out[2]) // out: 0xc00001a0e0 0xc00001a0e8 0xc00001a0f0
}

func goroutineTest() {
	list := []int{1, 2, 3}
	ch := make(chan int)
	for _, v := range list {
		go func(ch chan int) {
			ch <- v
			fmt.Printf("v:%d, &v:%p\n", v, &v)
		}(ch)
		<- ch
	}
/*
   v:1, &v:0xc0000140d8
   v:2, &v:0xc0000140d8
   v:3, &v:0xc0000140d8
*/
}
