package main

func myFunction(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	myFunction(66, 77)
}

// 使用下面命令执行：
// go tool compile -S -N -l 02-gotool.go