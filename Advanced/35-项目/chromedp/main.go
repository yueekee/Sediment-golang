package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	demo "github.com/liankui/Sediment-golang/Advanced/35-项目/chromedp/demo-检查资源占用情况"
)

// 使用 JSONP 向不同域的服务器请求数据。如果查询参数存在回调，则将回调添加到响应体中。
// 添加一个callback的入参后，会将callback对应的参数返回出来
func main() {
	app := gin.Default()

	// pprof router
	pprof.Register(app)

	app.POST("/report", demo.GetPDF) // html to pdf


	// 监听并在 0.0.0.0:8080 上启动服务
	app.Run(":8082")
}
