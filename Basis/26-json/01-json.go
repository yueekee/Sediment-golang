package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//s := `{"content":"\u5f88\u597d","friends":[{"content":"\u5356\u706b\u67f4\u7684\u5b89\u59ae","targetId":43},{"content":"asd","targetId":53}]}`
	//var m = make(map[string]interface{})

	s := `[{"content":"\u4e13\u680f666666","friends":[{"content":"\u5356\u706b\u67f4\u7684\u5b89\u59ae","targetId":43},
{"content":"asd","targetId":53}]},{"content":"\u4e13\u680f66666",
"friends":[{"content":"\u5356\u706b\u67f4\u7684\u5b89\u59ae","targetId":43},{"content":"asd","targetId":53}]}]`

	var m []map[string]interface{}
	json.Unmarshal([]byte(s), &m)
	fmt.Printf("m:%+v\n", m)
}
