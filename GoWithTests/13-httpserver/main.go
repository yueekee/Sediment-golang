package main

import (
	"fmt"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	fmt.Println("--start")
	serv := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", serv); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

// 用 curl http://localhost:5000/players/Pepper 获取玩家得分
