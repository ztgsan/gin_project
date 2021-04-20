package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Handlerfunc
func indexHandler(context *gin.Context)  {
	fmt.Println("index")
	context.JSON(http.StatusOK, gin.H{
		"message": "eee",
	})
}

// 中间件  m1
func m1(c *gin.Context) {
	fmt.Println("m1")
	start := time.Now()

	//c.Next()	 // 调用后续的处理函数
	c.Abort()	 // 组织后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost %v\n", cost)
	fmt.Println("m1 out!")

}

func loginAuth(doCheck bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		if doCheck {

		}else {

		}
	}

}

func main()  {

	r := gin.Default()
	r.GET("/index", m1, indexHandler)


	r.Use(m1)   // 全局注册中间件函数

	r.GET("/shop", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "shop",
		})
	})

	r.Run(":8888")


}
