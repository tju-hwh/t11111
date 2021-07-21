package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	r.GET("/login", func(c *gin.Context) {
		// 给客户端设置cookie
		//  maxAge int, 单位为秒
		// path,cookie所在目录
		// domain string,域名
		//   secure 是否智能通过https访问
		// httpOnly bool  是否允许别人通过js获取自己的cookie
		c.SetCookie("abc","123",60,"/","localhsot",false,true)
		c.String(200,"loginsuccess")
	})
	r.GET("/home",AuthoerMiddleWare(), func(c *gin.Context) {
		c.JSON(200,gin.H{"data":"home"})
	})
	r.Run(":8000")
}

func AuthoerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context){
		//获取客户端cookie并验证
		fmt.Println(c.Cookie("abc"))
		if cookie, err := c.Cookie("abc");err==nil{
			fmt.Println("有cookie")
			if cookie=="123"{
				c.Next()
				return
			}
		}
		//返回错误
		c.JSON(http.StatusUnauthorized,gin.H{"error":"err"})
		//若验证不通过，不再调用后续的函数处理
		//中止阻止调用挂起的处理程序。请注意，这不会停止当前处理程序。
		//假设您有一个验证当前请求是否已被授权的授权中间件。
		//如果授权失败（例如：密码不匹配），则调用Abort以确保剩余的处理程序
		//对于此请求，不调用。
		c.Abort()
		return
	}
}
