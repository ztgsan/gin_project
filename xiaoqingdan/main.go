package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

// 链接数据库
func initMysql()(err error){
	dns := "root:zhangtianguang@(127.0.0.1:3306)/xiaoqingdan?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open("mysql", dns)
	if err != nil {
		return err
	}
	return DB.DB().Ping()	 // 测试连通性
}

func main()  {

	// 连接数据库
	err := initMysql()
	if err != nil {
		panic(err)
	}

	defer DB.Close()		// 程序退出，关闭数据库

	// 模型绑定
	DB.AutoMigrate(&Todo{})		// 表名 todos

	r := gin.Default()

	// 指定html路径
	r.LoadHTMLGlob("templates/*")

	// 指定静态文件的路径
	r.Static("/static", "static")


	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})


	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(context *gin.Context) {
			var todo Todo
			context.BindJSON(&todo)
			//err := DB.Create(&todo)
			if err = DB.Create(&todo).Error; err != nil{
				context.JSON(http.StatusOK, gin.H{"message": err.Error()})
			}else {
				context.JSON(http.StatusOK, todo)
			}
		})

		// 查看所有
		v1Group.GET("/todo", func(context *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil{
				context.JSON(http.StatusOK, gin.H{"err": err.Error()})
			}else {
				context.JSON(http.StatusOK, todoList)
			}
		})

		// 查看某一个
		v1Group.GET("/todo/:id", func(context *gin.Context) {

		})

		// 删除
		v1Group.DELETE("/todo/:id", func(context *gin.Context) {
			id, ok := context.Params.Get("id")
			if !ok {
				context.JSON(http.StatusOK, gin.H{"err": "无效的ID"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil{
				context.JSON(http.StatusOK, gin.H{"err": err.Error()})
			}else {
				context.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
			}
		})

		// 修改
		v1Group.PUT("/todo/:id", func(context *gin.Context) {
			id, ok := context.Params.Get("id")
			if !ok {
				context.JSON(http.StatusOK, gin.H{"err": "无效的ID"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"err": err.Error()})
				return
			}
			context.ShouldBind(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"err": err.Error()})
				return
			}else {
				context.JSON(http.StatusOK, todo)
			}

		})

		// 完成

	}

	
	r.Run(":9999")
	
}
