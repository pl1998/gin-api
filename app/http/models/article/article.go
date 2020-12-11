package article

import (
	"goproject/pkg/model"
	"strconv"
)
type BaseModel struct {
	ID uint64
}

type Articles struct {
	ID          int64
	Title, Body string
}

//查询所有文章

func GetAll() ([]Articles, error) {

	var articles []Articles

	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles,nil
}

func (a BaseModel) GetStringID() string {
	return Uint64ToString(a.ID)
}

func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}