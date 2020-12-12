package routes

import (
	"github.com/gin-gonic/gin"
	Auth "goproject/app/http/auth/controller"
	"goproject/app/http/middleware"
)
var router *gin.Engine

//注册api

func RegisterApiRoutes(router *gin.Engine) {

	//api分组以及中间件

	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {    //类型断言
	//	v.RegisterValidation("loginvalidate",validate.LoginValidate) //注册调用tag和自定义验证器
	//}
	authc := new(Auth.AuthorizationController)

	router.POST("/api/login", authc.Login)
	router.POST("/api/register", authc.Register)

	api := router.Group("/api").Use(middleware.Auth())
	{
		api.GET("/users",authc.Users)
	}

}
