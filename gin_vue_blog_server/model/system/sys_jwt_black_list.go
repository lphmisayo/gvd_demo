package system

import (
	"gin_vue_blog_server/model"
)

type JwtBlackList struct {
	model.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
