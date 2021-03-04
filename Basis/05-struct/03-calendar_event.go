package main

import (
	"fmt"
	"github.com/liankui/Sediment-golang/Basis/05-struct/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year())
	err = event.SetTitle("sadf11")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
}
