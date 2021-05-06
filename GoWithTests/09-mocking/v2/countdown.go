package main

import (
	"fmt"
	"io"
	"time"
	"os"
)

/*
所谓迭代是指：确保我们采取最小的步骤让软件可用。
尽你所能拆分需求是一项很重要的技能，这样你就能拥有可以工作的软件。
缓慢的测试会破坏开发人员的生产力。

改进：
测试顺序的问题。
*/

const (
	finalWord = "Go!"
	countdownStart = 3
	write = "write"
	sleep = "sleep"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// CountdownOperationsSpy 同时实现了 io.writer 和 Sleeper，把每一次调用记录到 slice。
// 在这个测试中，我们只关心操作的顺序，所以只需要记录操作的代名词组成的列表就足够了。
type ConutdownOperationSpy struct {
	Calls []string
}

func (s *ConutdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *ConutdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}


type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{500 * time.Millisecond}
	Countdown(os.Stdout, sleeper)
}