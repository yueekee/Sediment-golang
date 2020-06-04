package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	// url encode
	v := url.Values{}
	v.Add("msg", "http://localhost:9999/&a=12")
	body := v.Encode()
	fmt.Println("s:", v.Get("msg"))
	//fmt.Println(v)
	fmt.Println("body:", body)	// http%3a%2f%2flocalhost%3a9999%2f%26a%3d12
	fmt.Println("trim:", strings.TrimLeft(body, "msg="))

	// url decode
	m, _ := url.ParseQuery(body)
	fmt.Println("m:", m)
}
