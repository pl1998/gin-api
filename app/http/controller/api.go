package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goproject/app/log"
	"net/http"
)

type Result struct {
	Code int64
	Msg string
	Data map[string]interface{}
}


type Articles struct {
	ID          int64
	Title, Body string
}

func Test(c *gin.Context) {
	Jstr := `
{"userId":"32f33eaa-f155-11e9-8a2a-00163e0e857e","passwordStrength":1,"nickName":"\u6f5c\u54f2","sex":"\u672a\u77e5","phone":"13217025359","email":null,"province":"\u5c71\u897f\u7701","city":"\u9633\u6cc9\u5e02","user_type":0,"country":"\u4e2d\u56fd","apple_id":"\u5df2\u7ed1\u5b9a","iconUrl":"https:\/\/accountdev.24e.co\/icon\/20201014\/20a41fb4960cffe37238e8938429c669.","languageType":"zn-ch","userType":"","wx_open_id":"\u672a\u7ed1\u5b9a","user_role":"\u666e\u901a\u7528\u6237","expiration_time":null,"cid":"d88904d2d50941c0b5895842c061b58d","langUrl":"https:\/\/img.24e.co\/lang\/zn-cn.json"}
`
	var person map[string]interface{}
	if err := json.Unmarshal([]byte(Jstr), &person); err != nil {
		log.LogError(err)
	}
	result := Result{Code: 200,Msg: "success",Data: person}

	c.JSON(http.StatusOK,result)
}


