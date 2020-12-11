package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册web路由
func RegisterWebRoutes(router *gin.Engine) {
	//模板路径
	router.LoadHTMLGlob("./templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":"golang",
		})
	})
}

