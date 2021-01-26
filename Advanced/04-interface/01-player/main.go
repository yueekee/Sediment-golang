package main

import "github.com/liankui/Sediment-golang/Advanced/04-interface/01-player/gadget"

type Player interface {	// 把接口定义在调用的包中更灵活
	Play(string)
	Stop()
}

//func PlayList(device gadget.TapePlayer, songs []string) {
func PlayList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"tomorrow", "only human", "sunday limit"}
	var player Player = gadget.TapePlayer{}
	PlayList(player, mixtape)
	player = gadget.TapeRecorder{} // 支持两种结构体了
	PlayList(player, mixtape)
}
