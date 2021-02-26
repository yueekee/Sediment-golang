package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type PCLecturer struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	AvatarURL  string    `json:"avatar_url"`
	Title      string    `json:"title"`
	Introduce  string    `json:"introduce"`
	//CreateTime time.Time `json:"create_time"`
	Type       int       `json:"type"`
	//UpdateTime utils.JsonTime  //`json:"update_time"`
}

func main() {
	s := `[
           {
               "id": 70,
               "name": "测试",
               "avatar_url": "https://statictest.chuanghehui.com/pc-teacher/1602573156527.png",
               "title": "",
               "introduce": "",
               "create_time": "2020-10-13 15:12:42",
               "type": 0
           },
           {
               "id": 69,
               "name": "w杜继南%￥",
               "avatar_url": "https://static.chuanghehui.com/pc-teacher/1589183842006.png",
               "title": "原美赞臣全国培训经理原美赞臣全",
               "introduce": "美国AACTP机构授权认证讲师、原美赞臣公司全国培训经理美国AACTP机构授权认证讲师、原美赞臣公司全国培训经理美国AACTP机构授权认证讲师、原美赞臣公司全国培训经理美国AACTP机构授权认证讲师、原美赞臣公司全国培训经理美国AACTP机构授权认证讲师、原美赞臣公司全国培训经理美国AACTP机构授权认证讲师、原美赞臣公司全国培训经理美国AACTP机构授权认证讲师、原美赞臣公司全国培训经理美国AA",
               "create_time": "2020-05-11 15:57:47",
               "type": 4
           }
       ]`

	//s := `{
    //    "lecturer_list":
    //        {
    //            "id": 70,
    //            "name": "测试",
    //            "avatar_url": "https://statictest.chuanghehui.com/pc-teacher/1602573156527.png",
    //            "title": "",
    //            "introduce": "",
    //            "create_time": "2020-10-13 15:12:42",
    //            "type": 0
    //        }
    //}`

	var m []PCLecturer
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("m:%#v\n", m)
	// m:[map[content:专栏666666 friends:[map[content:卖火柴的安妮 targetId:43]
	//map[content:asd targetId:53]]] map[content:专栏66666 friends:[map[content:卖火柴的安妮 targetnt:asd targetId:53]]]]
}
