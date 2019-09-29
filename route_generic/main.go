package main

import "github.com/gin-gonic/gin"

func main(){
	r:=gin.Default()
	r.GET("/user/*action", func(c *gin.Context) { // 所有/user开始的都会打到这里
		c.String(200, "hello world")
	})
	r.Run()
}
