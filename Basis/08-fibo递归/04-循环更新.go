package main

import "fmt"

func main() {
	p := make([]*Param, 0)
	p = append(p, &Param{ID: 1, ParentID: 0, Extra: nil})
	p = append(p, &Param{ID: 2, ParentID: 1, Extra: nil})
	p = append(p, &Param{ID: 3, ParentID: 1, Extra: nil})
	p = append(p, &Param{ID: 4, ParentID: 3, Extra: nil})

	fmt.Println(p)
	param := getParam(p)
	fmt.Println(param)
}

type Param struct {
	ID       int      `json:"id"`
	ParentID int      `json:"parent_id"`
	Extra    []*Param `json:"extra"`
}

// 假设p内的元素都是按ID、parentID从小到大的顺序
func getParam(p []*Param) []*Param {
	if p == nil {
		return nil
	}
	p1 := make([]*Param, 0)
	for i := 0; i < len(p); i++ {
		if p[i].ParentID == 0 {
			p1 = append(p1, p[i])
		}

	}

	return p
}

/* todo 预期结果，未实现
{
    "ID":1,
    "parent_id":0,
    "extra":[
        {
            "ID":2,
            "parent_id":1,
            "extra":null
        },
        {
            "ID":3,
            "parent_id":1,
            "extra":{
                "ID":4,
                "parent_id":3,
                "extra":null
            }
        }
    ]
}
*/