package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//1 创建路由
	r:=gin.Default()
	//2 绑定路由规则，执行的函数
	//gin.context 封装了request、和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"helloWorld")
	})
	//3 监听端口 默认8080
	//Run（“里面不指定端口号的话默认8080”）
	r.Run(":8080")
}
