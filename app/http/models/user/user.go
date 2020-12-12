package user

import (
	"goproject/pkg/model"
)

type Users struct {
	ID,TokenExpirationTime,CreateTime int64
	Email string  `valid:"email"`
	Password string  `valid:"password"`
	Token,Avatar,Username string
	OauthType int
	PasswordComfirm string ` gorm:"-" valid:"password_comfirm"`
}


//查询用户信息

func GetUsers(email string) (Users,error) {

	var user Users

	if err := model.DB.Where("email =?",email).First(&user).Error; err != nil {
		return user,err
	}

	return user,nil
}


