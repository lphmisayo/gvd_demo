package common

import "gin_vue_blog_server/model"

type User struct {
	model.Model
	Email     string `json:"email"`
	NickName  string `json:"nickName"`
	Avatar    string `json:"avatar"`
	Introduce string `json:"introduce"`
}
