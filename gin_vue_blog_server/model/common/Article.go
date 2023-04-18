package common

import "gin_vue_blog_server/model"

type Article struct {
	model.Model
	Title       string `json:"title" gorm:"type:varchar(100)"`
	Content     string `json:"content" gorm:"type:longtext"`
	Description string `json:"description" gorm:"type:varchar(200)"`
	Type        int8   `json:"type" gorm:"type:tinyint"`
	CategoryId  int    `json:"category_id" gorm:"type:bigint"`
	OriginUrl   string `json:"origin_url" gorm:"type:varchar(100)"`
	Status      int8   `json:"status" gorm:"type:tinyint"`
	AuthorId    int    `json:"author_id" gorm:"type:int"`
	ImageId     int    `json:"image_id" gorm:"type:int"`
	Stars       int    `json:"stars" gorm:"type:int"`
	Likes       int    `json:"likes" gorm:"type:int"`

	Author   User      `json:"author" gorm:"-"`
	Image    Image     `json:"image" gorm:"-"`
	Comments []Comment `json:"comments" gorm:"-"`
}
