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

	// 建表
	//db.AutoMigrate(&User{})

	//u1 := User{Name: "fd", Age: 18}
	//u2 := User{Name: "zhang", Age: 28}

	// 创建数据
	//db.Create(&u1)
	//db.Create(&u2)

	// 更新
	//var user User
	//db.First(&user)
	//user.Name = "我"
	//user.Age = 100
	//db.Save(&user)		// 所有字段全部更新

	//var user User
	//db.First(&user)
	//db.Model(&user).Update("name", "小玩意")		// 更新某个字段

	var user User
	m1 := map[string]interface{}{
		"name": "李文州",
		"age": 19,
		"active": true,
	}
	db.Model(&user).Updates(m1)

	db.Model(&user).Select("age").Updates(m1)	// 只更新age字段
	db.Model(&user).Select("active").Updates(m1)	// 排除active字段




}

