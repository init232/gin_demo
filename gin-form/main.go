package main

import "github.com/gin-gonic/gin"

//gin form表单
func main() {
	r := gin.Default()
	r.GET("/login", func(ctx *gin.Context) {
		ctx.GetPostForm("")
	})
	r.Run(":9090")
}
