package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age int64
	Active bool
}

func main() {

	db, err := gorm.Open("mysql", "root:zhangtianguang@(127.0.0.1:3306)/db1?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 删除
	//var user User
	//user.ID = 7
	//db.Debug().Delete(user)

	db.Where("name=?", "jinzhu").Delete(User{})
	db.Delete(User{}, "age=?", 18)

}


