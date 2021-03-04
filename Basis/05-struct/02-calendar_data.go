package main

import (
	"fmt"
	calendar "github.com/liankui/Sediment-golang/Basis/05-struct/calendar"
	"log"
)

// 总结：
// 将程序中的数据隐藏在一部分代码中且对另一部分不可见 的方法成为封装。
// 它可以用来防止无效数据。也可以修改程序代码的封装部分，不用担心其他代码的访问，而为它们不可直接访问。
//
// （这样在包内是可以直接访问未导出数据，但是对包外的，未导出的字段是隐藏的，这时不能意外地将值更改为无效的值）
func main() {
	date := calendar.Date{}	// 其中的字段为小写，不能通过赋值的方式输入值了
	err := date.SetYear(2000) // 只能通过setter方法更新，程序被保护，避免意外输入无效数据。
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(1)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(23)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
