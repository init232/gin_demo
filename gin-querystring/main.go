package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(ctx *gin.Context) {
		//获取浏览器发送的请求携带的query string参数
		//name := ctx.Query("query")
		//name := ctx.DefaultQuery("query", "somebody") //取不到就使用默认值
		name, ok := ctx.GetQuery("query") //根据取不到第二个参数就返回false
		if !ok {
			fmt.Println("取不到")
			name = "somebody"
		}
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"name": name,
			},
		)
	})
	r.Run(":9090")
}
