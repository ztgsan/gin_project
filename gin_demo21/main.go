package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type User struct {
	ID int64
	Name string `gorm:"default:'天光'"`		// 默认值
	Age int64
}

func main()  {

	db, err := gorm.Open("mysql", "root:zhangtianguang@(127.0.0.1:3306)/db1?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	//u1 := User{Age: 44}
	//db.Create(&u1)

	u := User{Name: "qimi", Age: 68}
	//fmt.Println(db.NewRecord(&u))	// 查询是不是一条新的记录
	db.Debug().Create(&u)

	//fmt.Println(db.NewRecord(&u))

}

