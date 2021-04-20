package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func main()  {
	r := gin.Default()
	r.GET("/web", func(context *gin.Context) {
		// 获取浏览器发请求鞋带的 query string 参数

		// 1.
		//name := context.Query("query")

		// 2. 取不到就用默认值
		name := context.DefaultQuery("query", "somebody")

		// 3. 取不到返回bool值
		name, ok := context.GetQuery("query")
		if !ok{
			name = "somebody"
		}

		fmt.Println(name)

		context.JSON(200, gin.H{
			"name": name,
		})

	})

	r.Run(":9999")


}
