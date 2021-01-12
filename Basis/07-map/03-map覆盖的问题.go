package main

import "fmt"

func main() {
	//resp := map[string]interface{}{
	//	"group":           1,
	//	"group_name":      "groupName",
	//	"group_commit_by": 1,
	//}

	resp := map[string]interface{}{}
	resp["group"] = 1

	resp = map[string]interface{}{
		"is_stop": 1,
	}

	fmt.Println(resp)
}
