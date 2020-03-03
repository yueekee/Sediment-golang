package main

/*
	1.defer需要在函数结束前执行
	2.函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数

*/

func main() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) { // t作用域为整个函数，defer在return前执行，返回4
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int { // t作用域不包括defer内的函数，返回1
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) { // t作用域为整个函数，返回3
	defer func() {
		t += i
	}()
	return 2
}
