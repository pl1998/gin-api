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
type AuthorizationController struct {}
// 用户登录
func (*AuthorizationController) Login(c *gin.Context) {
	var (
		email    = c.PostForm("email")
		password = c.PostForm("password")
	)
	user := userModel.Users{
		Email:    email,
		Password: password,
	}
	errs := requests.ValidateLoginForm(user)

	if len(errs) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 500,
			"msg":  errs,
		})
	} else {
		result :=model.DB.Select("password,id").Where("email=?",email).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 500,
				"msg":  "账号不存在",
			})
		} else {
			user_bool := helpler.ComparePasswords(user.Password,password)
			if user_bool {
				//颁发token
				generateToken(c,user)
			}else {
				c.JSON(http.StatusOK, map[string]interface{}{
					"code": 500,
					"msg":  "账号不存在",
				})
			}
		}
	}
}

func (*AuthorizationController) Users(c *gin.Context) {
	//通过该方法映射到map上
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	user := userModel.Users{}

	result := model.DB.Where("id =?",claims.ID).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 500,
			"msg":  "用户信息不存在",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":200,
		"msg":"获取用户信息成功",
		"data": map[string]interface{}{
			"email":user.Email,
			"username":user.Username,
			"avatar":user.Avatar,
			"create_time":time.Unix(user.CreateTime,0).Format("2006-01-02 15:04:05"),
		},
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
		register_time := time.Now().Unix()
		username := fmt.Sprintf("im_%v",strconv.FormatInt(register_time,10))
		//加密
		user := userModel.Users{
			Email:      c.PostForm("email"),
			Password:    helpler.HashAndSalt(c.PostForm("password")),
			Username:   username,
			CreateTime: register_time,
		}
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

	fmt.Println(expiration_time)

	fmt.Println(user.ID)
	j :=&jwt.JWT{
		[]byte(sign_key),
	}
	claims := jwt.CustomClaims{strconv.FormatInt(user.ID,10),jwtgo.StandardClaims{
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
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":200,
			"msg":"登录成功",
			"data": data,
		})
		return
	}

}


