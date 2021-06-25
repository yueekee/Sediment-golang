package main

import (
	"github.com/liankui/Sediment-golang/Advanced/35-项目/uerb-zap/logger"
)

func main() {
	logger.Info("在线周边poi信息查看消费失败，pid :", 1)
	logger.Fatal("在线周边poi信息查看消费失败，pid :", 1)
}
