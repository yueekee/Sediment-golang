package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Fatal(err)
	}
	fmt.Println(string(marshal))

	s := `[{"content":"\u4e13\u680f666666","friends":[{"content":"\u5356\u706b\u67f4\u7684\u5b89\u59ae","targetId":43},
{"content":"asd","targetId":53}]},{"content":"\u4e13\u680f66666",
"friends":[{"content":"\u5356\u706b\u67f4\u7684\u5b89\u59ae","targetId":43},{"content":"asd","targetId":53}]}]`

	var m []map[string]interface{}
	json.Unmarshal([]byte(s), &m)
	fmt.Printf("m:%+v\n", m)
	// m:[map[content:专栏666666 friends:[map[content:卖火柴的安妮 targetId:43]
	//map[content:asd targetId:53]]] map[content:专栏66666 friends:[map[content:卖火柴的安妮 targetnt:asd targetId:53]]]]
}
