package controller

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	userModel "goproject/app/http/models/user"
	"goproject/app/http/requests"
	jwt "goproject/pkg/auth"
	"goproject/pkg/config"
	"goproject/pkg/helpler"
	"goproject/pkg/model"
	"net/http"
	"strconv"
	"time"
)



type AuthorizationController struct {
}

// 用户登录
func (*AuthorizationController) Login(c *gin.Context) {

	create_time := time.Now()
	fmt.Println(create_time)
	user := userModel.Users{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	errs := requests.ValidateLoginForm(user)

	if len(errs) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 500,
			"msg":  errs,
		})
	} else {

		c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "success",
		})
	}
}

func (*AuthorizationController) Users(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "用户信息接口",
	})
}

// 用户注册
func (*AuthorizationController) Register(c *gin.Context) {

	user := userModel.Users{
		Email:           c.PostForm("email"),
		Password:        c.PostForm("password"),
		PasswordComfirm: c.PostForm("password_comfirm"),
	}
	errs := requests.RegisterValidateForm(user)

	if len(errs) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 500,
			"msg":  errs,
		})
	} else {
		//加密
		pwd := helpler.GetPwd(c.PostForm("password"))

		user := userModel.Users{
			Email:      c.PostForm("email"),
			Password:    helpler.HashAndSalt(pwd),
			Username:   "im_" + strconv.Itoa(time.Now().Second()),
			CreateTime: time.Now().Unix(),
		}

		fmt.Println(user)

		result := model.DB.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code": 500,
				"msg":  "用户注册失败",
			})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":200,
				"msg":"用户注册成功!",
			})
		}

	}
}

//获取token

func generateToken(c *gin.Context,user userModel.Users) {


	sign_key            := config.GetString("app.jwt.sign_key")
	expiration_time     := config.GetInt("app.jwt.expiration_time")

	j :=&jwt.JWT{
		[]byte(sign_key),
	}

	ids := fmt.Sprintf(`"json:%s"`,user.ID)

	claims := jwt.CustomClaims{ids,jwtgo.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,
		ExpiresAt: time.Now().Unix() + int64(expiration_time),
		Issuer: sign_key,
	}}


	token,err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusForbidden, map[string]interface{}{
			"code":403,
			"msg":"jwt token颁发失败--",
		})

		return

	} else {

		data := map[string]interface{}{
			"token":token,
			"user":user,
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"code":200,
			"msg":"登录成功",
			"data": data,
		})
		return
	}

}


