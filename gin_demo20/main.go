package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)


type User struct {
	gorm.Model
	Name string
	Age sql.NullInt64
	Birthday *time.Time
	Email string `gorm:"type:varchar(100);unique_index"`
	Role string `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num int `gorm:"AUTO_INCREMENT"`
	Address string `gorm:"index:addr"`   // 给本字段创建名为addr的多音
	IgnoreMe int `gorm:"-"`

}


type Animal struct {
	AinmalId int64 `gorm:"primary_key"`
	Name string
	Age int64
}

func (a Animal) TableName()string {
	return "qimi"
}

func main()  {

	db, err := gorm.Open("mysql", "root:zhangtianguang@(127.0.0.1:3306)/db1?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.SingularTable(true) // 禁用复数

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	// 使用User结构体创建名叫 小王子 的表
	db.Table("xiaowangzi").CreateTable(&User{})

}


