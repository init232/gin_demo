package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

// gin 上传
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(ctx *gin.Context) {
		//从请求中读取文件
		file, err := ctx.FormFile("f1")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//dst := fmt.Sprintf("./%s", file.Filename)   //通过拼接方式
			dst := path.Join(file.Filename) //path处理路径相关
			//将读取的文件保存在本地（服务端本地）
			ctx.SaveUploadedFile(file, dst)
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		}

	})
	r.Run(":9090")
}
