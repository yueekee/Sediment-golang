package calendar

import "errors"

/*总结：结构体的内部字段的验证和更新
1.如果是不想暴露的字段使用小写
2.通过绑定的方法对字段进行内部验证或更新
*/

type Date struct {
	year  int
	month int
	day   int
}

// setter方法是用来将值设置给struct字段或者变量的方法。
// getter方法是获取struct字段或者变量值的方法。
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Day() int {
	return d.day
}
