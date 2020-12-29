package main

import "fmt"

func main() {
	p := make([]*Param, 0)
	p = append(p, &Param{ID: 1, ParentID: 0, Extra: nil})

	c := make([]*Param, 0)
	c = append(c, &Param{ID: 2, ParentID: 1, Extra: nil})
	c = append(c, &Param{ID: 3, ParentID: 1, Extra: nil})
	c = append(c, &Param{ID: 4, ParentID: 3, Extra: nil})

	fmt.Println(p)
	fmt.Println(c)
	param := getParam(p, c)
	fmt.Println(param)
}

type Param struct {
	ID       int      `json:"id"`
	ParentID int      `json:"parent_id"`
	Extra    []*Param `json:"extra"`
}

func getParam(p, c []*Param) []*Param {
label:
	for i := 0; i < len(p); i++ {
		for i2 := 0; i2 < len(c); i2++ {
			if c[i2].ParentID == p[i].ID {
				p[i].Extra = append(p[i].Extra, c[i2])
				if i2 == 0 {
					c = c[1:]
				} else if i2 == len(c)-1 {
					c = c[:i2]
				} else {
					c = append(c[:i2], c[i2+1:]...)
				}
				i2--
			}
		}
	}
	if len(c) > 0 {
		fmt.Println("p1", p)
		for i := 0; i < len(p); i++ {
			if p[i].Extra != nil {
				p = p[i].Extra
				fmt.Println("p1", p)
				goto label
			}
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