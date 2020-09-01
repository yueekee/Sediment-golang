package main

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func main() {
	a, b := 2, 3
	max := If(a > b, a, b).(int)
	println(max)
}
