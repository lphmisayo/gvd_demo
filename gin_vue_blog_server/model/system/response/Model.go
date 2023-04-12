package response

import "time"

type Model struct {
	ID       uint      `json:"ID"  gorm:"primary_key;auto_increment"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}
