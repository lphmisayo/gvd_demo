package article

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model"
	"gin_vue_blog_server/model/common"
	"gin_vue_blog_server/model/response"
	"gorm.io/gorm"
)

type ArticleListService struct {
}

type Option struct {
	model.PageInfo
	Debug bool
}

func (*ArticleListService) GetAuthorArticlesSimpleInfo(model common.Article, option Option) (list []response.ArticleSimpleInfo, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLogConfig.MysqlLog})
	}

	if option.Sort == "" {
		option.Sort = "create_at desc" //默认时间倒序
	}

	count = DB.Select("author_id").Find(&list).RowsAffected

	offest := (option.Page - 1) * option.Limit
	if offest < 0 {
		offest = 0
	}
	err = DB.Limit(option.Limit).Offset(option.Limit).Order(option.Sort).Find(&list).Error
	return list, count, err
}
func (*ArticleListService) GetArticles(model common.Article, option Option) (list []common.Article, count int64, err error) {
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
