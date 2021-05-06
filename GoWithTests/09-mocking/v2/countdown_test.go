package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T){
		buffer := &bytes.Buffer{}
		Countdown(buffer, &ConutdownOperationSpy{})
	
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("sleep after every print", func(t *testing.T){
		spySleeperPrinter := &ConutdownOperationSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			sleep, 
			write,
			sleep, 
			write,
			sleep, 
			write,
			sleep, 
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("want calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})
}