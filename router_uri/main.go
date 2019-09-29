package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.GET("/:name/:id", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": context.Param("name"),
			"id": context.Param("id"),
		})
	})
	r.Run()
}
