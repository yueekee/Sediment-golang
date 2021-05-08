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
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T){
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)
	
		defer slowServer.Close()
		defer fastServer.Close()
	
		slowURL := slowServer.URL
		fastURL := fastServer.URL
	
		want := fastURL
		got, err := Racer(slowURL, fastURL)
	
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T){
		slowServer := makeDelayServer(11 * time.Millisecond)
		defer slowServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, slowServer.URL, 25 * time.Millisecond)
	
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
