package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()

	r.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "get",
		})
	})

	r.POST("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "post",
		})
	})

	r.Any("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "any",
		})
		switch context.Request.Method {
		case http.MethodPost:
			context.JSON(http.StatusOK, gin.H{"method": "any"})
		case http.MethodGet:
			context.JSON(http.StatusOK, gin.H{"method": "get"})
		}
	})

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"msg": "没有"})
	})

	// 视频的首页和详情页
	// 商城的首页和详情页

	//r.GET("/video/index", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	//})
	//
	//r.GET("/shop/index", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	//})

	// 路由组的概念
	videoFroup := r.Group("/video")
	{
		videoFroup.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoFroup.GET("/xxxxx", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "/video/xxxxx"})
		})
		videoFroup.GET("/ooooo", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "/video/ooooo"})
		})
	}

	r.Run(":8888")

}
