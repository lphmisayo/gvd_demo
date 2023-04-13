package model

import (
	"gorm.io/gorm/logger"
	"time"
)

type Model struct {
	ID       uint      `json:"ID" gorm:"primary_key;auto_increment"`
	CreateAt time.Time `json:"CreateAt"`
	UpdateAt time.Time `json:"UpdateAt"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type MysqlLogSt struct {
	MysqlLog logger.Interface
	Debug    bool
}
