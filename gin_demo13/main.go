package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// 获取请求的path (URI) 参数

func main()  {

	r := gin.Default()

	r.GET("/user/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
		})
	})


	r.GET("/blog/:year/:month", func(context *gin.Context) {
		year := context.Param("year")
		month := context.Param("month")
		context.JSON(http.StatusOK, gin.H{
			"year": year,
			"month": month,
		})
	})

	r.Run(":9999")

}
