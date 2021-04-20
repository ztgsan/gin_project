package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

// 链接数据库
func InitMysql()(err error){
	dns := "root:zhangtianguang@(127.0.0.1:3306)/xiaoqingdan?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open("mysql", dns)
	if err != nil {
		return err
	}
	return DB.DB().Ping()	 // 测试连通性
}

func Close() {
	DB.Close()
}