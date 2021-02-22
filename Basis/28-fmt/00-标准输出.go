package main

import (
	"errors"
	"fmt"
)

/*
Printf 带格式的打印，2个特性：
	1.格式化动词
	2.值的宽度
例如：%7.3f
%后面加数字就是对值进行宽度限制，如果小于宽度则用空格填充
其中7-整个数的最小宽度，3-小数点后的宽度，f-格式化动词的类型
动词:
	%f 浮点数
	%d 十进制整数
	%s 字符串
	%t 布尔值
	%v 任何值（根据所提供的值的类型选择适当的格式）
	%+v 任何值，包含字段在内的实例的完整信息
	%#v 任何值，按其在Go程序代码中显示的格式进行格式化
	%T 所提供值的类型
	%% 一个完完全全的百分号
*/

func main() {

	fmt.Printf("%f\n", 3.1415)
	fmt.Printf("%d\n", 12)
	fmt.Printf("%s\n", "hello")
	fmt.Printf("%t\n", true)
	fmt.Printf("%v\n", ErrDemo)	// this is a error
	fmt.Printf("%v %v %v\n", 1.2, "\t", true)	// 1.2      true
	fmt.Printf("%#v %#v %#v\n", 1.2, "\t", true)	// 1.2 "\t" true

	fmt.Println("-------------------")
	var d = Demo{ID: 11, Name: "eric"}
	fmt.Printf("%v\n", d)                        // {11 eric}
	fmt.Printf("%+v\n", d)                       // {ID:11 Name:eric}
	fmt.Printf("%#v\n", d)                       // main.Demo{ID:11, Name:"eric"}
	fmt.Printf("%T %T %T\n", 3.1415, "\t", true) // float64 string bool
	fmt.Printf("%%\n")                           // %

	fmt.Println("-------------------")
	// %后面加数字就是对值进行宽度限制，如果小于宽度则用空格填充
	// %7.3f 其中7-整个数的最小宽度，3-小数点后的宽度，f-格式化动词的类型
	fmt.Printf("%%7.3f: %7.3f\n",  12.3456)	// %7.3f:  12.346
	fmt.Printf("%%7.2f: %7.2f\n",  12.3456)	// %7.2f:   12.35
	fmt.Printf("%%7.1f: %7.1f\n",  12.3456)	// %7.1f:    12.3
	fmt.Printf("%%.1f: %.1f\n",  12.3456)	// %.1f: 12.3     四舍五入到一位小数，无填充
	fmt.Printf("%%.2f: %.2f\n",  12.3456)	// %.2f: 12.35
}

var ErrDemo = errors.New("this is a error")

type Demo struct {
	ID   int
	Name string
}
