package main

import "fmt"

/*总结：
第二个defer因为f()未被定义，导致panic错误，随后执行第一个defer，异常被recover()，程序恢复正常
有个迷惑的点，就r和n的关系。最后return的是n+1，但是出参为r，这里表示r=n+1即3+1=4，然后在执行第一个defer中r+=n即4+3=7
*/

func f(n int) (r int) {
	defer func() {
		fmt.Println("-------defer1,r:", r)
		fmt.Println("-------defer1,n:", n)
		r += n
		recover()
		fmt.Println("-------r:", r)
	}()

	var f func()
	defer f()

	fmt.Println("-------")
	f = func() {
		r += 2
		fmt.Println("-------r2", r)
	}

	fmt.Println("-------n", n)
	return n + 1
	/*
	-------
	-------n 3
	-------defer1,r: 4
	-------defer1,n: 3
	-------r: 7
	7
	*/

}

func main() {
	fmt.Println(f(3))
}