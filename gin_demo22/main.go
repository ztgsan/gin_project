package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age int64
}

func main()  {

	db, err := gorm.Open("mysql", "root:zhangtianguang@(127.0.0.1:3306)/db1?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 建表
	//db.AutoMigrate(&User{})

	//u1 := User{Name: "fd", Age: 18}
	//u2 := User{Name: "zhang", Age: 28}

	// 创建数据
	//db.Create(&u1)
	//db.Create(&u2)

	// 查询
	var user User
	db.First(&user)		// 主键必须是int
	fmt.Println(user)


	var users []User
	db.Debug().Find(&users)
	fmt.Println(users)


	var user1 User
	db.FirstOrInit(&user1, User{Name: "小王子"})




}