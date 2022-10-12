package main

import "github.com/gin-gonic/gin"

// gin 路由和路由组
func main() {
	r := gin.Default()
	r.GET("/index", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})
	//Norouter
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "访问的页面不存在",
		})
	})
	//视频首页和详情
	//r.GET("/video/index", func(ctx *gin.Context) {
	//	ctx.String(200, "video/index")
	//})
	//r.GET("/video/index", func(ctx *gin.Context) {
	//	ctx.String(200, "video/1")
	//})
	//r.GET("/video/index", func(ctx *gin.Context) {
	//	ctx.String(200, "video/2")
	//})
	// 路由组感念
	videGroup := r.Group("/video")
	{
		videGroup.GET("/index", func(ctx *gin.Context) {
			ctx.String(200, "index")
		})
		videGroup.GET("/1", func(ctx *gin.Context) {
			ctx.String(200, "1")
		})
	}

	////商城首页和详情
	//r.GET("/shop/index", func(ctx *gin.Context) {
	//	ctx.String(200, "shop/index")
	//})
	//r.GET("/shop/index", func(ctx *gin.Context) {
	//	ctx.String(200, "shop/1")
	//})
	//r.GET("/shop/index", func(ctx *gin.Context) {
	//	ctx.String(200, "shop/2")
	//})
	r.Run(":9090")
}
