package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type Student struct {
	ID int
	Name string
}

func main() {
	var s1 = Student{
		ID: 1,
		Name: "eric",
	}
	marshal, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))

	// 使用这种更快
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, err := json.Marshal(&s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))

	// -----------
	s := `[{"content":"\u4e13\u680f666666","friends":[{"content":"\u5356\u706b\u67f4\u7684\u5b89\u59ae","targetId":43},
{"content":"asd","targetId":53}]},{"content":"\u4e13\u680f66666",
"friends":[{"content":"\u5356\u706b\u67f4\u7684\u5b89\u59ae","targetId":43},{"content":"asd","targetId":53}]}]`

	var m []map[string]interface{}
	json.Unmarshal([]byte(s), &m)
	fmt.Printf("m:%+v\n", m)

	json.Unmarshal([]byte(s), &m)
	fmt.Printf("m2:%+v\n", m)
}
