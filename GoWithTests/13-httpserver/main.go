package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	serv := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", serv); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
