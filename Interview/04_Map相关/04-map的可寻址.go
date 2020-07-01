package main

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
	m["zhangsan"].Age = 23
	m["zhangsan"].GrowUp()

	p := m["zhangsan"]
	p.Age = 23
	p.GrowUp()
	m["zhangsan"] = p
}
func main() {
	m := map[string]*Person{
		"zhangsan": &Person{Age: 20},
	}

	m["zhangsan"].Age = 23
	m["zhangsan"].GrowUp()
}