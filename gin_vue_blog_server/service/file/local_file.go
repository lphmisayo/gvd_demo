package file

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model"
	"gorm.io/gorm"
)

type LocalFileService struct {
}

type Option struct {
	model.PageInfo
	Debug bool
}

func GetFileList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLogConfig.MysqlLog})
	}

	if option.Sort == "" {
		option.Sort = "create_at desc" //默认时间倒序
	}

	count = DB.Select("id").Find(&list).RowsAffected

	offest := (option.Page - 1) * option.Limit
	if offest < 0 {
		offest = 0
	}
	err = DB.Limit(option.Limit).Offset(option.Limit).Order(option.Sort).Find(&list).Error
	return list, count, err
}
