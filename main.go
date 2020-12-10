package main

import (
	"github.com/gin-gonic/gin"
	"goproject/bootstrap"
	"goproject/routes"
)

//入口文件 入口函数
func main() {
	router := gin.Default()
	//加载数据库
	bootstrap.SetupDB()
	//注册路由
	routes.RegisterApiRoutes(router)
	//监听端口
	router.Run(":8080")
}
