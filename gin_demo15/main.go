package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

	r := gin.Default()

	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(context *gin.Context) {
		// 读取文件
		f ,err := context.FormFile("f1")
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
		}else {
			context.SaveUploadedFile()
		}

	})

	r.Run(":9999")


}
