package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 获取URL路径参数
func main() {
	r := gin.Default()
	r.GET("/:name/:age", func(ctx *gin.Context) {
		//获取路径参数
		name := ctx.Param("name")
		age := ctx.Param("age")
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(ctx *gin.Context) {
		year := ctx.Param("year")
		month := ctx.Param("month")
		ctx.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	r.Run(":9090")
}
