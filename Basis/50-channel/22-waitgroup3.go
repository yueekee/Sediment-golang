package main

import (
	"errors"
	"fmt"
	"sync"
)

type errInfo struct {
	err error
	id  int
}

var errCh = make(chan *errInfo)
var waitGroup sync.WaitGroup

func main() {
	go rangeCh()

	waitGroup.Add(9)
	for i := 1; i < 10; i++ {
		go readDB(&waitGroup, i)
	}
	waitGroup.Wait()

}

func rangeCh() {
	for err := range errCh {
		// 有错误可以立马打印出来
		fmt.Println(err)
	}
}

func readDB(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	err := read_db(i)
	if err != nil {
		errCh <- &errInfo{
			err: err,
			id:  i,
		}
	}
}

func read_db(i int) error {
	fmt.Println("num:", i)
	if i == 3 {
		return errors.New("3 error")
	}
	return nil
}
