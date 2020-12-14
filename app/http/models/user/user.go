package user

import (
	"goproject/pkg/model"
	"time"
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


func (a Users) CreateTimeAt() string{
	return time.Unix(a.CreateTime,0).Format("2006-01-02 15:04:05")
}

func (a Users) GetAvatar() string {
	if a.Avatar =="" {
		return "https://together.24e.co/qsj_default_avatar.png"
	}

	return a.Avatar
}


