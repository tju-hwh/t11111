package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		//获取客户端是否携带cookie


		c.SetCookie("key_cookie", "value_cookie", 60, "/","localhost",false,true)
		cookie11, err := c.Cookie("key_cookie")
		if err != nil {
			fmt.Println(" error: ", err)
			//return
		}

		fmt.Printf("cookie的值是： %s\n",cookie11)
		c.String(http.StatusOK,cookie11)
	})
	r.Run(":8080")
}
