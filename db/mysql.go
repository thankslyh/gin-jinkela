package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var mysqlDB *gorm.DB

func init() {
	log.Println("init mysql")
	dsn := "root:123456@tcp(127.0.0.1:3307)/jinkela_schema?charset=utf8&parseTime=True&loc=Local"
	mysqlDB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println("连接错误", err.Error())
	//	return
	//}
}

func GetMysqlDB() *gorm.DB  {
	return mysqlDB
}
