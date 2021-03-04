package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*总结：切片的指针问题
1.使用%p打印slice，打印的是slice指向底层数组的第一个元素地址
2.如果slice长度为0，打印的是其它地址
*/

func main() {
	months := []string{1: "January", 2: "two"}
	fmt.Printf("%p,\n", &months)    // 0xc0000a6020,
	fmt.Printf("%p,\n", months)     // 0xc0000a6040,
	fmt.Printf("%p,\n", &months[0]) // 0xc0000a6040,
	fmt.Printf("%p,\n", &months[1]) // 0xc0000a6050,

	num := []int{}
	fmt.Printf("%p,\n", &num)    // 0xc0000a6060,
	fmt.Printf("%p,\n", num)     // 0x119e428,

	header := (*reflect.SliceHeader)(unsafe.Pointer(&months))
	fmt.Printf("%p,\n", header)          // 0xc0000a6020, months 变量地址 (SliceHeader)
	fmt.Printf("0x%x,\n", header.Data)   // 0xc0000a6040, 从 months 变量地址处取一个指针变量的长度当做指针打印 (SliceHeader 开头也正好是 uintptr)
	fmt.Printf("0x%x,\n", header.Data+0) // 0xc0000a6040, 打印 slice 指向的底层数组第一个元素地址
}
