package model

import "time"

type Model struct {
	ID       uint64    `json:"ID" gorm:"primary_key;auto_increment"`
	CreateAt time.Time `json:"CreateAt"`
	UpdateAt time.Time `json:"UpdateAt"`
}
