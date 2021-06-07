package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// c.Next和c.Abort的用法

func main() {
	r := gin.Default()

	d := "play"
	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := r.Group("/admin").Use(func(c *gin.Context) {
		//设计一种简单的中间件进行判断客户端传输的数据是否有效
		if d == "play" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "成功进入相关的业务",
			})
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			fmt.Println("结束权限访问功能")
			c.Abort() //终止下面的函数执行
		}
	})
	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		c.JSON(http.StatusOK, gin.H{"user": "gcg", "secret": "123456"})
	})
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}