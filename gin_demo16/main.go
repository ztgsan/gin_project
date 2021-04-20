package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main()  {

	r := gin.Default()

	r.GET("/index", func(context *gin.Context) {
		//context.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})
		// 跳转到百度
		context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")			// 重定向
	})
	
	r.GET("/a", func(context *gin.Context) {
		context.Request.URL.Path = "/b"
		r.HandleContext(context)
	})

	r.GET("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":9999")

}
