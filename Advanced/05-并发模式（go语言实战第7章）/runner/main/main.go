package main

import (
	"github.com/liankui/Sediment-golang/Advanced/05-并发模式（go语言实战第7章）/runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	r := runner.NewRunner(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1) // 如果执行超时，用错误码1终止
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2) // 如果执行中断信号，用错误码2终止
		}
	}

	log.Println("Process ended")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
