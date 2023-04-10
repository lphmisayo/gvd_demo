package common

import (
	"gin_vue_blog_server/model"
	ctype "gin_vue_blog_server/model/cType"
)

type Image struct {
	model.Model
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `json:"name" gorm:"type:varchar(100)"`
	ImageType ctype.ImageType `json:"image_type" gorm:"type:int8"`
}
