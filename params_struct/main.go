package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Name string	`form:"name"`
	Address string	`form:"address"`
	Birthday time.Time	`form:"birthday" time_format:"2006-01-02"`
}

func main(){
	r:=gin.Default()
	r.GET("/test", testing)
	r.POST("/test", testing)
	r.Run()
}

func testing(c *gin.Context){
	var person Person
	// 根据content-type来做不同的binding操作
	if	err:=c.ShouldBind(&person); err == nil {
		c.String(http.StatusOK, "%v", person)
	} else {
		c.String(http.StatusOK, "person bind error %v", err)
	}
}
