package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main(){
	// 如果需要把日志不在控制台输出而是输出到文件
	f,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r:=gin.New()	// 默认会使用这两个中间件Logger, Recovery logger可以在有请求是，打出一条日志
	r.Use(gin.Logger(), gin.Recovery())	// 加上recovery之后，如果成功发生崩溃，可以阻止错误继续上报，从而程序崩溃
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(200, "%s", name)
	})
	r.Run()
}
