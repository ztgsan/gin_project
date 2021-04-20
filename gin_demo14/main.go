package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main()  {

	r := gin.Default()
	
	r.GET("/user", func(context *gin.Context) {
		// 参数绑定
		var user UserInfo
		err := context.ShouldBind(&user)
		if err != nil {
			context.JSON(400, gin.H{
				"err": err.Error(),
			})
		}else {
			fmt.Printf("%#v\n", user)
			context.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}

		//username := context.Query("username")
		//password := context.Query("password")
		//
		//user := UserInfo{
		//	username,
		//	password,
		//}
		fmt.Printf("")
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.POST("/form", func(context *gin.Context) {
		var user UserInfo
		err := context.ShouldBind(&user)
		if err != nil {
			context.JSON(400, gin.H{
				"err": err.Error(),
			})
		}else {
			fmt.Printf("%#v\n", user)
			context.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.Run(":9999")

}
