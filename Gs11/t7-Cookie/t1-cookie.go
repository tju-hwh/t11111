package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	r.GET("cookie", func(c *gin.Context) {
		//获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {

			// 给客户端设置cookie
			//  maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			//   secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			cookie="not set"
			c.SetCookie("key_cookie", "value_cookie", 60, "/","localhost",false,true)

		}
		fmt.Printf("cookie的值是： %s\n", cookie)
		c.String(http.StatusOK,"succ")
	})
	r.Run(":8080")
}
