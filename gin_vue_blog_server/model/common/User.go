package common

import "gin_vue_blog_server/model"

type User struct {
	model.Model
	UserName  string `json:"userName" gorm:"column:username"`
	PassWord  string `json:"passWord" gorm:"column:password"`
	Email     string `json:"email"`
	NickName  string `json:"nickName" gorm:"column:nickname"`
	Avatar    string `json:"avatar" `
	Introduce string `json:"introduce"`
}
