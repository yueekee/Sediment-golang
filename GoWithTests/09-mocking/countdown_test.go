package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	SpySleeper := &SpySleeper{}

	Countdown(buffer, SpySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if SpySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", SpySleeper.Calls)
	}
}