package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"log"
)

/*总结：
1、浅拷贝：对基本数据类型进行值传递，对引用数据类型进行引用传递般的拷贝，此为浅拷贝。
2、深拷贝：对基本数据类型进行值传递，对引用数据类型，创建一个新的对象，并复制其内容，此为深拷贝。
浅拷贝改变一个对象其中引用类型的地址，会影响到另一个对象。
*/

type Dog struct {
	age  int
	name string
	tags []string
}

func main() {
	swallowCopy()
	deepCopy()
	deepcopy2()
}

func swallowCopy() {
	roger := Dog{5, "Roger", []string{"a", "b"}}
	mydog := roger
	mydog.tags[0] = "c"
	fmt.Printf("%#v\n", roger)	// main.Dog{age:5, name:"Roger", tags:[]string{"c", "b"}}
	fmt.Printf("%#v\n", mydog)	// main.Dog{age:5, name:"Roger", tags:[]string{"c", "b"}}
}

// 人工深拷贝
func deepCopy() {
	roger := Dog{5, "Roger", []string{"a", "b"}}
	mydog := roger

	mydog.tags = make([]string, len(roger.tags))
	copy(mydog.tags, roger.tags)

	mydog.tags[0] = "c"
	fmt.Println(roger)	// {5 Roger [a b]}
	fmt.Println(mydog)	// {5 Roger [c b]}
}

// 使用其它包,todo 好像不行。。。。
func deepcopy2() {
	roger := Dog{5, "Roger", []string{"a", "b"}}
	mydog := Dog{}
	err := copier.Copy(&mydog, &roger)
	if err != nil {
		log.Fatal(err)
	}

	mydog.tags = append(mydog.tags, "c")
	fmt.Println(roger)	// {5 Roger [a b]}
	fmt.Println(mydog)	// {0  [c]}
}
