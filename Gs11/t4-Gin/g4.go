package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types:=c.DefaultPostForm("type","post")
		username:=c.PostForm("username")
		password:=c.PostForm("userpassword")
		c.String(http.StatusOK,fmt.Sprintf("u:%s p:%s t:%s",username,password,types))
	})
	r.Run()
}