package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 会从该链接中下载文件，命名为filename中的名字
func main() {
	router := gin.Default()
	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://cn.bing.com/th?id=OHR.CortezJacks_ZH-CN1619906832_1920x1080.jpg&rf=LaDigue_1920x1080.jpg&pid=hp")
		if err != nil || response.StatusCode != http.StatusOK {
			fmt.Println("err:", err)
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run(":8080")
}