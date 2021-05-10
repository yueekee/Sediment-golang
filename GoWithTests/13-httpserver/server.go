package main

import (
	"fmt"
	"net/http"
)

/*
net/http 的 ResponseWriter 也实现了 io Writer，所以我们可以用 fmt.Fprint 发送字符串来作为 HTTP 应答.
你应该养成优先让编译通过再让测试通过的编程习惯。
*/

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, r)
	case http.MethodPost:
		p.processWin(w)
	}
	
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/Players/"):]
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, GetPlayerScore(player))
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

// 我们把得分计算从 PlayerServer 移到函数 GetPlayerScore 中，这就是使用接口重构的正确方法。
func GetPlayerScore(name string) string {
	if name == "Pepper" {
        return "20"
    }

    if name == "Floyd" {
        return "10"
    }

	return ""
}