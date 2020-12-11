package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goproject/app/http/models/article"
	"goproject/app/log"
	"goproject/bootstrap"
	"net/http"
)

type Result struct {
	Code int64
	Msg string
	Data map[string]string
}


type Articles struct {
	ID          int64
	Title, Body string
}

func Test(c *gin.Context) {

	articles,err :=article.GetAll()

	path :=bootstrap.GetCurrentPath()
	fmt.Print(path)
	if err != nil {
		log.LogError(err)
	}else {
		fmt.Println(articles)
		c.JSON(http.StatusOK,articles)
	}
}


