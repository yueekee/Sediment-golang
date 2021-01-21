package main

import "github.com/liankui/Sediment-golang/Advanced/04-interface/gadget"

func PlayList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapeRecorder{}
	mixtape := []string{"tomorrow", "only human", "sunday limit"}
	PlayList(player, mixtape)	// TapeRecorder无法使用PlayList函数
}
