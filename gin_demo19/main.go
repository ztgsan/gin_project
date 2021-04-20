package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo --> 表
type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

func main()  {
	db, err := gorm.Open("mysql", "root:zhangtianguang@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	// 创建表
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	//u1 := UserInfo{
	//	1,
	//	"天光",
	//	"男",
	//	"跑步",
	//}
	//db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u)	// 查询表中第一条数据
	fmt.Println(u)

	// 更行
	//db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)


}
