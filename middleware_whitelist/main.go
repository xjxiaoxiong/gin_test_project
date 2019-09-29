package main

import "github.com/gin-gonic/gin"


func IPAuthMiddleware() gin.HandlerFunc{
	return func(context *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag:=false
		clientIP := context.ClientIP();
		for _,host := range ipList {
			if clientIP==host {
				flag = true
				break
			}
		}
		if !flag {
			context.String(401, "%s, not in whitelist", clientIP)
			context.Abort()
		}
	}
}

func main(){
	r:=gin.Default()
	r.Use(IPAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})
	r.Run()
}
