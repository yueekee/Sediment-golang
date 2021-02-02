package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
		}(&wg)
	}
	wg.Wait()
}

// GODEBUG=schedtrace=1000 ./01-GODEBUG.go
// GODEBUG=scheddetail=1,scheddetail=1000 ./01-GODEBUG
