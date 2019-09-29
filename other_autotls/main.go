package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})
	// 自动化证书配置： 调用证书下载的包,然后就是https的过程：1、生成本地密钥，2、然后用这个密钥去证书颁发机构获取私钥，3、进行本地私钥验证，验证成功把私钥保存，下次再有请求，用这个私钥加密
	autotls.Run(r, "www.itpp.tk") // 运行需要真实域名
}

