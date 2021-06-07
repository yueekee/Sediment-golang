package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 提供 unicode 实体
	// 输出 : {"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	// 输出 : {"html":"<b>Hello, world!</b>"}
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
