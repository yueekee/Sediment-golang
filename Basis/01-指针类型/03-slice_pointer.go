package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	months := []string{1: "January"}
	fmt.Printf("%p,\n", &months)
	fmt.Printf("%p,\n", months)
	fmt.Printf("%p,\n", &months[0])

	header := (*reflect.SliceHeader)(unsafe.Pointer(&months))
	fmt.Printf("%p,\n", header)          // months 变量地址 (SliceHeader)
	fmt.Printf("0x%x,\n", header.Data)   // 从 months 变量地址处取一个指针变量的长度当做指针打印 (SliceHeader 开头也正好是 uintptr)
	fmt.Printf("0x%x,\n", header.Data+0) // 打印 slice 指向的底层数组第一个元素地址
}
