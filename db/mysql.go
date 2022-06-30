package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jinkela/setting"
)

var mysqlDB *gorm.DB

func SetMysqlUp() {
	fmt.Println("init mysql=", setting.DatabaseSetting)
	dsn := "root:123456@tcp(" + setting.DatabaseSetting.Host + ")/jinkela?charset=utf8&parseTime=True&loc=Local"
	mysqlDB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println("连接错误", err.Error())
	//	return
	//}
}

func GetMysqlDB() *gorm.DB  {
	return mysqlDB
}
