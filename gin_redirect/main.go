package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 请求重定向
func main() {
	r := gin.Default()
	r.GET("/index", func(ctx *gin.Context) {
		ctx.Redirect(301, "https://www.google.com")
	})
	//路由重定向
	r.GET("/a", func(ctx *gin.Context) {
		//跳转到/b对应的路由处理函数
		ctx.Request.URL.Path = "/b" //把请求的URL修改
		r.HandleContext(ctx)        //继续后续的处理
	})
	r.GET("/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "b",
		})
	})
	r.Run(":9090")
}
