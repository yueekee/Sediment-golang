package main

import "github.com/gin-gonic/gin"

// 通常，JSON 使用 unicode 替换特殊 HTML 字符，例如 < 变为 \ u003c。如果要按字面对这些字符进行编码，则可以使用 PureJSON。
// Go 1.6 及更低版本无法使用此功能。
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
