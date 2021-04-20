package routers

import (
	"gin_project/xiaoqingdan/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 指定html路径
	r.LoadHTMLGlob("templates/*")

	// 指定静态文件的路径
	r.Static("/static", "static")

	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有
		v1Group.GET("/todo", controller.GetTodoList)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteOneTodo)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateOneTodo)

		// 完成

	}
	return r
}