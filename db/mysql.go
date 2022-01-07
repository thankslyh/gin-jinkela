package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var mysqlDB *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/jinkela_schema?charset=utf8&parseTime=True&loc=Local"
	mysqlDB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println("连接错误", err.Error())
	//	return
	//}
}

func GetMysqlDB() *gorm.DB  {
	return mysqlDB
}
