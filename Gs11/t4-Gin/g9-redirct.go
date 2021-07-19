package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	c := gin.Default()
	v1:=c.Group("/v1")
	{
		v1.GET("/t",readhh)
		v1.GET("/v",jujuju)
	}
	c.Run()
}

func readhh(ctx *gin.Context)  {
	ctx.String(http.StatusOK,fmt.Sprintf("good"))
}

func jujuju(ctx *gin.Context)  {
	ctx.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
}