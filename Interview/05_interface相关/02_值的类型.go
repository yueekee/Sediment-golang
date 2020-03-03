package main

func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}

// 解释：type只能使用在interface
// func GetValue() int { 返回的是int类型时，i.(type)是已知类型，不符合规则
func GetValue() interface{} {
	return 1
}
