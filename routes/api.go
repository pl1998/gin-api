package routes

import (
	"github.com/gin-gonic/gin"
	"goproject/app/http/controller"
	"goproject/app/http/middleware"
	"net/http"
)
var router *gin.Engine

//注册api
func RegisterApiRoutes(router *gin.Engine) {

	//模板
	router.GET("/", controller.List)

	//api分组以及中间件
	api := router.Group("/api").Use(middleware.Auth())
	{
		api.GET("/test", func(c *gin.Context) {
			example := c.MustGet("example").(string)
			c.JSON(http.StatusOK, map[string]interface{}{
				"token": example,
			})
		})
		api.GET("/json", controller.Test)
	}

}
