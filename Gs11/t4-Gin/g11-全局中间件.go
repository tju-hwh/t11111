package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	r:=gin.Default()
	r.Use(MiddleWare())
	{
		r.GET("/ce", func(c *gin.Context) {
			value, _ := c.Get("name")
			fmt.Println("name: ",value)
			//页面接收
			c.JSON(200,gin.H{"name":value})
		})
	}
	r.Run()
}

func MiddleWare() gin.HandlerFunc{
	return func(g *gin.Context) {
		t:=time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		g.Set("name","李")
		g.Next()
		// 中间件执行完后续的一些事情
		status:=g.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
