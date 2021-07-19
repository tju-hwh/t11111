package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	r.GET("/user", func(c *gin.Context) {
		name:=c.DefaultQuery("name","不传的话就打印这一段默认值")
		c.String(http.StatusOK,fmt.Sprintf("hello %s",name))
	})

	r.Run()
}
