package common

import "gin_vue_blog_server/model"

type Category struct {
	model.Model
	Name string `json:"name" gorm:"type:varchar(50);not null"`

	Articles []Article `json:"articles" gorm:"-"`
}
