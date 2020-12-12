package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goproject/app/http/models/user"
)

// 用户登录效验
func ValidateLoginForm(data user.Users) map[string][]string{

	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email"},
		"password":         []string{"required", "min:6"},
	}
	messages := govalidator.MapData{

		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
		},
	}


	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}


	errs := govalidator.New(opts).ValidateStruct()

	return errs
}
