package main

import (
	"fmt"
	"time"
	"unsafe"
)

/*总结：
struct{}不占据任何内存空间
场景：
1.实现方法接收者；
2.实现集合类型；
3.实现空通道。

还有作为context value的key也算一个使用场景
*/

// 1.实现方法接收者；(也便于未来添加公共字段)
type T struct {
}

func (t *T) Call()  {
}

func Scene1() {
	var t T
	t.Call()
}

// 2.实现集合类型；(只需要用key，不需要用value)
type Set map[string]struct{}

func (s Set) Append(k string) {
	s[k] = struct{}{}
}

func (s Set) Remove(k string) {
	delete(s, k)
}

func (s Set) Exist(k string) bool {
	_, ok := s[k]
	return ok
}

// 3.实现空通道。（可用于流转各类状态或是控制并发情况）
// 打印1111之后sleep一秒后再打印2222
func Scene3() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	fmt.Println("1111")
	<-ch
	fmt.Println("2222")
}

func Scene2() {
	set := Set{}
	set.Append("apple")
	set.Append("pineapple")
	set.Append("pear")
	fmt.Println(set)

	set.Remove("apple")
	fmt.Println(set)
	fmt.Println(set.Exist("pear"))
}


func main() {
	Scene3()
	var a int
	var b string
	var c bool
	var d [3]int32
	var e []string
	var f map[string]bool
	var g struct{}

	fmt.Println(
		unsafe.Sizeof(a),
		unsafe.Sizeof(b),
		unsafe.Sizeof(c),
		unsafe.Sizeof(d),
		unsafe.Sizeof(e),
		unsafe.Sizeof(f),
		unsafe.Sizeof(g),
	)
	// 8 16 1 12 24 8 0
}



// incomparable is a zero-width, non-comparable type. Adding it to a struct
// makes that struct also non-comparable, and generally doesn't add
// any size (as long as it's first).
type incomparable [0]func()
