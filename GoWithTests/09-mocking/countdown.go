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
*/

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

const (
	finalWord = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{500 * time.Millisecond}
	Countdown(os.Stdout, sleeper)
}