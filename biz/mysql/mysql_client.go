package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 单例mysql client
var dbCli *gorm.DB

func init() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/course_arragement?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	dbCli, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}