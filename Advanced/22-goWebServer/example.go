// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*
重点看下http.Request结构体部分和其中的r.URL结构体

*/

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/2", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// handler2 echoes the HTTP request.
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

/*
GET /2 HTTP/1.1
Header["Upgrade-Insecure-Requests"] = ["1"]
Header["User-Agent"] = ["Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36"]
Header["Sec-Fetch-User"] = ["?1"]
Header["Cookie"] = ["Phpstorm-5064742f=319dc684-3667-44b8-afb9-e3f7466f4109; UM_distinctid=17743463e05921-0b25c2ffefab48-163e655c-1fa400-17743463e06db0; CNZZDATA1279445308=898498664-1611740677-null%7C1611740677; Hm_lvt_c368b8fe55b947e24b94c99f1bf8463b=1611741413; Goland-f5f3fb11=76da86f1-43d3-418c-95dc-bb488f2c1d3c"]
Header["Sec-Ch-Ua-Mobile"] = ["?0"]
Header["Sec-Fetch-Mode"] = ["navigate"]
Header["Sec-Ch-Ua"] = ["\"Google Chrome\";v=\"89\", \"Chromium\";v=\"89\", \";Not A Brand\";v=\"99\""]
Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng
Header["Sec-Fetch-Site"] = ["none"]
Header["Accept-Encoding"] = ["gzip, deflate, br"]
Header["Accept-Language"] = ["zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7"]
Header["Connection"] = ["keep-alive"]
Header["Sec-Fetch-Dest"] = ["document"]
Host = "localhost:8000"
RemoteAddr = "127.0.0.1:56135"
*/