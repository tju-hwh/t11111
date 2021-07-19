package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	v3:=r.Group("/v3")
	{
		v3.GET("/testha",testHandler)
		v3.GET("/testhb",testb)

	}

	if err:=r.Run();err!=nil{
		fmt.Printf("gg")
	}

}

func testHandler(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"message":"hello gin!",
	})

}

func testb(c *gin.Context)  {
	name:=c.DefaultQuery("name","linacan")
	c.String(http.StatusOK,fmt.Sprintf("2222%s",name))

}