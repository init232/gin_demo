package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// gin 中间件

// handlerfunc
func indexHandler(ctx *gin.Context) {
	ctx.String(200, "ok")
}

// 定义中间件
func m1(ctx *gin.Context) {
	fmt.Println("m1进来了")
	//记时
	start := time.Now()
	ctx.Next() //调用后续的处理函数
	//ctx.Abort()   //阻止调用后续的函数
	cost := time.Since(start)
	fmt.Printf("cosgt:%v\n", cost)
}

//func authMiddeware(ctx *gin.Context) {
//	//是否登录的判断
//	// if 是登录用户  c.next()
//	//否则  ctx.Abort()
//
//}

// 闭包 中间件
func authMiddeware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
func main() {
	r := gin.Default()
	r.Use(m1, authMiddeware())
	r.GET("/index", m1, indexHandler)
	r.GET("/user", func(ctx *gin.Context) {
		ctx.String(200, "user")
	})
	r.GET("/info", func(ctx *gin.Context) {
		ctx.String(200, "info")
	})
	r.Run()
}
