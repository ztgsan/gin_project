package controller

import (
	"gin_project/xiaoqingdan/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(context *gin.Context)  {
	context.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(context *gin.Context) {
	var todo models.Todo
	context.BindJSON(&todo)
	//err := DB.Create(&todo)

	err := models.CreateATodo(&todo)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}else {
		context.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(context *gin.Context) {
	todoList, err := models.GetAllTodos()
	if err != nil{
		context.JSON(http.StatusOK, gin.H{"err": err.Error()})
	}else {
		context.JSON(http.StatusOK, todoList)
	}
}

func DeleteOneTodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"err": "无效的ID"})
		return
	}
	err := models.DeleteATodo(id)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{"err": err.Error()})
	}else {
		context.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
	}
}

func UpdateOneTodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"err": "无效的ID"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	context.ShouldBind(&todo)

	err = models.UpdateATodo(todo)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}else {
		context.JSON(http.StatusOK, todo)
	}

}