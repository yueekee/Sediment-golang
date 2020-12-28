package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

const (
	taskLoad2        = 10
	goroutineNumber = 4
	PopNumber       = 4
)

type task struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func main() {
	var wg sync.WaitGroup
	channels := make(chan string, taskLoad2)
	for i := 0; i < goroutineNumber; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			select {
			case ch := <- channels:
				fmt.Println("+++ch:", ch)
				var t task
				err := json.Unmarshal([]byte(ch), &t)
				if err != nil {
					fmt.Println("err:", err)
					//return
				}
				time.Sleep(time.Second * 5)
				fmt.Println("--------------------------")
				fmt.Println("++++t.ID", t.ID)
			}
		}(&wg, i)
	}
	//tasks := []task{{1, "eric"}}
	tasks := []task{{1, "eric"}, {2, "yoly"}, {3, "loria"}}
	for _, task := range tasks {
		jsonTask, err := json.Marshal(task)
		if err != nil {
			fmt.Println("err1:", err)
		}
		channels <- string(jsonTask)
	}
	close(channels)
	wg.Wait()
}
