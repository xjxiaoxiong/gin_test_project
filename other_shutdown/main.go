package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	r:=gin.Default()
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(8*time.Second)
		c.String(200, "hello test\n")
	})
	srv:=&http.Server{
		Addr:":8085",
		Handler: r,
	}
	go func() {
		if err:=srv.ListenAndServe(); err!=nil && err!=http.ErrServerClosed{ // 并且服务器没有关闭
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	// 退出信号的捕获
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("看chan阻塞")
	<-quit	// chan进行阻塞
	fmt.Print("拿到信号了")
	log.Print("...shutdown server...")

	// 设定超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 真正关闭服务器
	if err:=srv.Shutdown(ctx); err!=nil {
		log.Fatal("server shutdown", err)
	}

	log.Println("server exiting")
}
