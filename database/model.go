package database

import (
	"goproject/app/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

type BaseModel struct {
	ID uint64
}


// 初始化 grom
func ConnectDB() *gorm.DB {

	var err error
	config := mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})
	DB,err = gorm.Open(config,&gorm.Config{

		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	log.LogError(err)
	return DB

}



