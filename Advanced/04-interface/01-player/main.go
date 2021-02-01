package main

import (
	"fmt"
	"github.com/liankui/Sediment-golang/Advanced/04-interface/01-player/gadget"
)

// 类型断言时使用第二参数ok避免panic的情况

type Player interface { // 把接口定义在调用的包中更灵活
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

func TryOut(player Player) {
	player.Play("Test Track")	// Playing2 Test Track todo 这里为什么会直接就是定位为TapeRecorder类型来
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder) // 使用类型断言来获得TapeRecorder
	if ok {
		recorder.Record()
	}

	p, ok := player.(gadget.TapePlayer)	// 断言，使用ok对回报panic的情况做处理。
	if ok {
		p.Stop()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func main() {
	mixtape := []string{"tomorrow", "only human", "sunday limit"}
	var player Player = gadget.TapePlayer{}
	PlayList(player, mixtape)
	player = gadget.TapeRecorder{} // 支持两种结构体了
	PlayList(player, mixtape)

	// --------------
	fmt.Println("-----------")
	TryOut(gadget.TapeRecorder{})
}
