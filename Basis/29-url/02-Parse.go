package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := "https://pay.chuanghehui.com/index/chh161604332827?inviter_id=123&name=eric"
	sourceUrl, _ := url.Parse(u)
	val, _ := url.ParseQuery(sourceUrl.RawQuery)
	fmt.Printf("sourceUrl:%#v\n", sourceUrl)
	fmt.Printf("val:%#v\n", val)

/*
   sourceUrl:&url.URL{Scheme:"https", Opaque:"", User:(*url.Userinfo)(nil), Host:"pay.chuanghehui.com",
Path:"/index/chh161604332827", RawPath:"", ForceQuery:false, RawQuery:"inviter_id=123&name=eric",
Fragment:"", RawFragment:""}

   val:url.Values{"inviter_id":[]string{"123"}, "name":[]string{"eric"}}
*/
}
