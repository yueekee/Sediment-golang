package main

import "fmt"

func main() {
	// 不覆盖

	//resp := map[string]interface{}{
	//	"group":           1,
	//	"group_name":      "groupName",
	//	"group_commit_by": 1,
	//}
	//
	//resp["test"] = 1
	//
	//fmt.Println(resp)	// map[group:1 group_commit_by:1 group_name:groupName test:1]

	// 覆盖

	resp := map[string]interface{}{}
	resp["group"] = 1

	resp = map[string]interface{}{
		"is_stop": 1,
	}

	fmt.Println(resp)	// map[is_stop:1]
}
