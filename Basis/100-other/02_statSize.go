package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	t1 := time.Now().UnixNano()
	//url := "https://pic.cnblogs.com/face/725163/20150226170205.png"
	//header,err := http.Head(url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("header:%+v\n", header)
	////fmt.Println(header.ContentLength)

	url2 := "http://wx.qlogo.cn/mmopen/kr9ShguI1uzOvkB15VXrdPWegpCW6OHd9sbdgPmicVTIMibT3bMd226KZb4ibv1XQDjV8lH0DdpjZRQ6vkO47zMibjU8cEQWGhrD/0"
	header2, err := http.Head(url2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("header2:%+v\n", header2.Header["X-Errno"])
	t2 := time.Now().UnixNano()
	fmt.Println("t2:", t2)
	fmt.Println("t1:", t1)
	fmt.Println("time:", t2 - t1)
	/*
	header:&{Status:200 OK StatusCode:200 Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Accept-Ranges:[bytes] Cache-Control:[max-age=600] Connection:[keep-alive] Content-Length:[3889] Content-Type:[image/png] Date:[Sun, 07 Jun 2020 06:52:18 GMT] Expires:[Sun, 07 Jun 2020 07:02:18 GMT] Last-Modified:[Fri, 24 Nov 2017 14:20:54 GMT] Server:[NWS_TCloud_S1] Strict-Transport-Security:[max-age=31536000] X-Cache-Lookup:[Hit From Disktank3 Hit From Inner Cluster] X-Daa-Tunnel:[hop_count=1] X-Nws-Log-Uuid:[1288b092-81aa-4d1e-ba04-b79bcf9d9834]] Body:{} ContentLength:3889 TransferEncoding:[] Close:false Uncompressed:false Trailer:map[] Request:0xc0000f4000 TLS:0xc0003138c0}
	header:&{Status:200 OK StatusCode:200 Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1
	Header:map[
	Cache-Control:[no-cache] Content-Length:[5093] Content-Type:[image/png]
	Last-Modified:[Mon, 01 Jan 1990 00:00:00 GMT] Server:[ImgHttp3.0.0] X-Bcheck:[0_0]
	X-Cpt:[filename=0] X-Errno:[-6101] X-Info:[notexist:-6101] X-Rtflag:[0]
	]
	Body:{} ContentLength:5093 TransferEncoding:[] Close:false Uncompressed:false Trailer:map[] Request:0xc0000f4100 TLS:<nil>}


	*/
}
