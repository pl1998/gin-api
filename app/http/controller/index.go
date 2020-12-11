package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {

}

var router *gin.Engine


func HomeIndex(c *gin.Context) () {
	c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"title": "home page",
	})
}