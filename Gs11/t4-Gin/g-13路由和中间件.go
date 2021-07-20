package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r:=gin.Default()
	r.Use(computeTime)
	g1:=r.Group("/g1")
	{
		g1.GET("/index",indexHandler)
		g1.GET("/home",homeHandler)
	}


	r.Run()
}

func computeTime(c *gin.Context){
	start:=time.Now()
	c.Next()
	since:=time.Since(start)
	str:=fmt.Sprintf("程序用时：%s",string(since))
	c.Set("status",str)

}

func indexHandler(c *gin.Context)  {
	time.Sleep(5*time.Second)
	//c.JSON("200",gin.H{"status",status}

}

func homeHandler(c *gin.Context)  {
	time.Sleep(2*time.Second)

}

