package main

/*总结：
f.bar不能使用 := 进行赋值
*/

type foo struct {
	bar int
}

func main() {
	var f foo
	//f.bar, tmp = 1, 2
	//fmt.Println(tmp)
	f.bar, _ = 1, 2
}
