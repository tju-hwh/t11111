package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Regist struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main(){
	r := gin.Default()
	r.GET("/:user/:password", func(c *gin.Context) {
		var regist Regist
		if err:=c.ShouldBindUri(&regist);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if regist.User!="user" || regist.Pssword!="admin"{
			c.JSON(http.StatusBadRequest,gin.H{"status":"200"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":"200"})
	})
	r.Run()
}


