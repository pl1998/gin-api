package main

import (
	"github.com/gin-gonic/gin"
	"goproject/bootstrap"
	"goproject/config"
	c "goproject/pkg/config"
	"goproject/routes"
)

func init() {
	//初始化加载一些配置文件
	config.Initialize()
}

//入口文件 入口函数
func main() {
	router := gin.Default()
	//加载数据连接池
	bootstrap.SetupDB()

	//注册路由
	routes.RegisterWebRoutes(router) //web路由
	routes.RegisterApiRoutes(router) //api路由

	//监听http端口
	router.Run(":"+c.GetString("app.port"))
}
