package main

import (
	"fmt"
	"log"
)

// 定义满足内建接口的类型来调用接口
// type error interface {
//	 Error() string
// }
// 同理还有stringer
// type Stringer interface {
// 	String() string
// }

type comedyError string
func (c comedyError) Error() string {	// comedyError满足error的接口
	return string(c)
}

// -------------
type OverheatError float64
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}
// 定义一个校验温度的函数
func checkTemperature(actual, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	var err error
	err = comedyError("111")	// 111
	fmt.Println(err)

	// ----------
	// 可以创建自己的满足error接口的类型并保存你所需要的信息
	var err2 error = checkTemperature(123.1, 110.0)
	if err2 != nil {
		log.Fatal(err2)	// 2021/02/01 23:20:15 Overheating by 13.10 degrees!
	}
}
