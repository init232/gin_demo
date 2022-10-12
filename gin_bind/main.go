package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 参数绑定
type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(ctx *gin.Context) {
		//username := ctx.Query("username")
		//password := ctx.Query("password")
		//u := UserInfo{
		//	username: username,
		//	password: password,
		//}
		var u UserInfo //声明一个userinfo类型的变量u
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		}
	})

	r.POST("/form", func(ctx *gin.Context) {
		var u UserInfo //声明一个userinfo类型的变量u
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		}
	})
	r.Run(":9090")
}
