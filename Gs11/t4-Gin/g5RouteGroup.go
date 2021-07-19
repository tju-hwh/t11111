package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	v1:=r.Group("/v1")
	{
		v1.GET("/login",login)
		v1.GET("/submit",submit)
	}

	v2:=r.Group("/v2")
	{
		v2.GET("/login",login)
		v2.GET("/submit",submit)
	}
	r.Run(":8000")
}

func login(c *gin.Context)  {
	name:=c.DefaultQuery("name","jack")
	c.String(http.StatusOK,fmt.Sprintf("hello  %s",name))
}

func submit(c *gin.Context)  {
	name:=c.DefaultQuery("name","lulu")
	c.String(http.StatusOK,fmt.Sprintf("hello  %s",name))
}