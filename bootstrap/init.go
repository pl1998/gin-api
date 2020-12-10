package bootstrap

import (
	"goproject/database"
	"time"
)



func SetupDB() {

	db := database.ConnectDB()

	sqlDB,_ :=db.DB()

	sqlDB.SetMaxOpenConns(50)
	//设置最大空闲数
	sqlDB.SetMaxIdleConns(25)
	//设置每个连接的超时时间
	sqlDB.SetConnMaxLifetime(5*time.Minute)
}


