package bootstrap

import (
	"goproject/pkg/config"
	"goproject/pkg/model"
	"goproject/pkg/redis"
	"time"
)



func SetupDB() {

	db := model.ConnectDB()

	sqlDB,_ :=db.DB()

	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	//设置最大空闲数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	//设置每个连接的超时时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)


	redis.InitClient()
}




