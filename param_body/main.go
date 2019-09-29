package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main(){
	r:=gin.Default()
	r.POST("/test", func(context *gin.Context) {
		bodyBytes, err :=  ioutil.ReadAll(context.Request.Body)
		if err!=nil {
			context.String(http.StatusBadRequest, err.Error())
			context.Abort()
		}

		// 前面已经读过一次，这里再次注入
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstName := context.PostForm("first_name")
		lastName := context.DefaultPostForm("last_name", "last_default_name")
		context.String(http.StatusOK, "%s,%s,%s", firstName, lastName, string(bodyBytes))
		//context.String(http.StatusOK,  string(bodyBytes))
	})
	r.Run()
}
