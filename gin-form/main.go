package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(ctx *gin.Context) {
		username := ctx.DefaultPostForm("username", "samebody")
		password := ctx.DefaultPostForm("password", "")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Name": username,
			"pass": password,
		})
	})
	r.Run(":9090")
}
