package main

import (
	"github.com/gin-gonic/gin"
	error2 "github.com/liankui/Sediment-golang/Advanced/34-gin/error/v1"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/err1", func(c *gin.Context) {
		error2.ResSuccess(c, "这是data")
	})

	r.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务
	// test
}
