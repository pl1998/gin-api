package routes

import (
	"github.com/gin-gonic/gin"
	"goproject/pkg/config"
	"net/http"

)

//注册web路由
func RegisterWebRoutes(router *gin.Engine) {
	//模板路径
	router.LoadHTMLGlob("./templates/*")
	router.GET("/", func(c *gin.Context) {

		title := config.Env("APP_NAME")

		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":title,
		})
	})
}

