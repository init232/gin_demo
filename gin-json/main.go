package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/map", func(ctx *gin.Context) {
		//方法一 map
		//data := map[string]interface{}{
		//	"name":    "tim",
		//	"age":     18,
		//	"address": "shanghai",
		//}
		data := gin.H{"name": "tim", "age": 20, "address": "shanghai"}
		ctx.JSON(http.StatusOK, data)
	})
	r.GET("/struct", func(ctx *gin.Context) {
		//方法二 结构体
		type msg struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		data := msg{
			"tom",
			20,
		}
		ctx.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
