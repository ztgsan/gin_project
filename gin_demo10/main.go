package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main()  {

	r := gin.Default()
	r.GET("/json", func(context *gin.Context) {
		// 方法1. map
		//data := map[string]interface{}{
		//	"name": "小王子",
		//	"message": "hello world",
		//	"age": 18,
		//}

		//方法2. gin.H
		data := gin.H{
			"name": "小王子",
			"message": "hello world",
			"age": 18,
		}

		context.JSON(http.StatusOK, data)
	})


	// 方法3. 结构体. 使用tag 定制化操作
	type msg struct {
		Name string `json:"name"`
		Message string `json:"message"`
		Age int `json:"age"`
	}

	r.GET("/another_json", func(context *gin.Context) {
		data := msg{
			"小王子",
			"hello golang",
			30,
		}
		context.JSON(http.StatusOK, data)
	})

	r.Run(":9999")

}