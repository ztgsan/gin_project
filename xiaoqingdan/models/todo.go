package models

import "gin_project/xiaoqingdan/dao"

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

// Todo 增删改查操作都放这

func CreateATodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil{
		return err
	}
	return
}


func GetAllTodos()(todoList []*Todo, err error){
	if err = dao.DB.Find(&todoList).Error; err != nil{
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error){
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil{
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo)(err error)  {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string)(err error)  {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}