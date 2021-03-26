package main

/*总结：
map的值是不可寻址的
可以使用临时变量或者声明指针类型进行处理
*/

type Person struct {
	Age int
}

func (p *Person) GrowUp() {
	p.Age++
}

func main1() {
	m := map[string]Person{
		"zhangsan": Person{Age: 20},
	}
	//m["zhangsan"].Age = 23	// 报错，不能直接更新
	//m["zhangsan"].GrowUp()	// 报错，指针函数无法执行

	p := m["zhangsan"]	// 方法1： 使用临时变量，更新赋值
	p.Age = 23
	p.GrowUp()
	m["zhangsan"] = p
}

func main() {
	m := map[string]*Person{	// 方法2：使用指针类型
		"zhangsan": &Person{Age: 20},
	}

	m["zhangsan"].Age = 23
	m["zhangsan"].GrowUp()
}