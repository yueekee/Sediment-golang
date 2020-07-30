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
	fmt.Println("s:", v.Get("msg")) // msg->http://localhost:9999/&a=12
	fmt.Println("body:", body)      // body: msg=http%3A%2F%2Flocalhost%3A9999%2F%26a%3D12

	trimLeft := strings.TrimLeft(body, "msg=")
	fmt.Println("trim:", trimLeft)  // trim: http%3A%2F%2Flocalhost%3A9999%2F%26a%3D12

	// url decode
	m, _ := url.ParseQuery(body)
	fmt.Println("m:", m["msg"][0])  // m: map[msg:[http://localhost:9999/&a=12]]
	// url decode 2
	unescape, _ := url.QueryUnescape(trimLeft)
	fmt.Println("un:", unescape)
}
