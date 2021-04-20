package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func main()  {

	r := gin.Default()
	r.LoadHTMLFiles("./login.html")
	
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(context *gin.Context) {
		username := context.PostForm("username")
		password, ok := context.GetPostForm("password")
		if !ok {
			password = "************"
		}

		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	
	r.Run(":9999")

}

