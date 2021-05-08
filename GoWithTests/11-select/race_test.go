package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*
在标准库中有一个 net/http/httptest 包，它可以让你轻易建立一个 HTTP 模拟服务器（mock HTTP server）。

httptest.NewServer 接受一个我们传入的 匿名函数 http.HandlerFunc。
http.HandlerFunc 是一个看起来类似这样的类型：type HandlerFunc func(ResponseWriter, *Request)。
*/

func TestRacer(t *testing.T) {
	slowServer := makeDelayServer(20 * time.Millisecond)
	fastServer := makeDelayServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got, _ := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T){
		slowServer := makeDelayServer(11 * time.Second)
		fastServer := makeDelayServer(12 * time.Second)
	
		defer slowServer.Close()
		defer fastServer.Close()
	
		_, err := Racer(slowServer.URL, fastServer.URL)
	
		if err != nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
