package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//// 方式一 最普通方式
//func sayhello(w http.ResponseWriter, r *http.Request) {
//	// 1. 直接返回
//	fmt.Fprint(w, "我是你大爷")
//	// 2. 获取用户输入内容返回
//	//var ret string
//	//fmt.Printf("请输入您想要返回的内容")
//	//fmt.Scanln(&ret)
//	//fmt.Fprint(w, ret)
//
//}
//
//func main() {
//	http.HandleFunc("/hello", sayhello)
//	http.ListenAndServe(":9999", nil)
//
//}


// 方式二, 从文件中读取要返回的内容
//func sayhello(w http.ResponseWriter, r *http.Request)  {
//
//	b, err := ioutil.ReadFile("./aaa.txt")
//	if err != nil {
//		fmt.Println("文件读取失败")
//	}
//	fmt.Fprint(w, string(b))
//
//}
//
//func main()  {
//	http.HandleFunc("/hello", sayhello)
//	http.ListenAndServe(":9999", nil)
//
//}


// 方式三, 返回json格式的数据

func sayhello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "我是OK",
	})

}

func main() {
	r := gin.Default()
	r.GET("hello, golang!", sayhello)

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "我是get请求",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "我是post请求",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "我是put请求",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "我是delete请求",
		})
	})

	r.Run(":9999")			// 默认是8080端口
}

