package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// 创建在 goroutine 中使用的副本
		cCp := c.Copy()
		go func() {
			// 用 time.Sleep() 模拟一个长任务。
			time.Sleep(5 * time.Second)

			// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
			log.Println("Done! in path " + cCp.Request.URL.Path)
			//cCp.JSON(200, gin.H{
			//	"mes": "123",
			//
			//})
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(5 * time.Second)

		// 因为没有使用 goroutine，不需要拷贝上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

func main2() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		go func() {   // 这儿开了协程，导致函数直接返回了，
			// 然后context 可能已经写了header，因为直接返回，
			// 所以context(非并发安全)可能认为的 响应数据为0，因此
			// content-length: 0
			// 但是后续又执行了 c.JSON() , body里面又有数据，
			// 因此客户端解析包出错了，才会报上面那个错误，
			// 正确的写法是这里不用go func
			c.JSON(200, gin.H{
				"message": "pong",
			})
		}()

	})
	r.Run()
}