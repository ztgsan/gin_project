package main

import (
	"gin_project/xiaoqingdan/dao"
	"gin_project/xiaoqingdan/models"
	"gin_project/xiaoqingdan/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main()  {

	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}

	defer dao.Close()		// 程序退出，关闭数据库

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})		// 表名 todos

	r := routers.SetupRouter()

	r.Run(":9999")
	
}
