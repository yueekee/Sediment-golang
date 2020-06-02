package main

import "fmt"

// 接口的定义

type Text interface {
	Pages(p int) int
	Words() int
	PageSize() int
}

type Book string

func (b *Book) Pages(p int) int {
	return p
}

func (b *Book) Words() int {
	return len(*b)
}

func (b *Book) PageSize() int {
	return len(*b)-1
}

type eBook struct {
	Book
	price int
}

func (e *eBook) Pages(p int) int {
	return p
}

func main() {
	var b Book
	b = "eric"
	fmt.Println(b.Pages(100))
	fmt.Println(b.Words())
	fmt.Println(b.PageSize())
}
