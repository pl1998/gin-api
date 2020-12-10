package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {

}

var router *gin.Engine

func List(c *gin.Context) ()  {

	router.LoadHTMLFiles("./../.././templates/index.tmpl")

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Posts",
	})
}