package common

import "gin_vue_blog_server/model"

type Tag struct {
	model.Model
	Name string `json:"name" gorm:"type:varchar(20);not null"`

	Articles []Article `json:"articles" gorm:"-"`
}
