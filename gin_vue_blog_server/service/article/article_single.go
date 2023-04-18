package article

import (
	"errors"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/common"
	"gin_vue_blog_server/model/request"
	"gorm.io/gorm"
)

type ArticleSingleService struct {
}

func (ArticleSingleService) GetArticleDetail(articleDetailReq request.ArticleDetailReq, debug bool) (*common.Article, error) {
	DB := global.DB
	if debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLogConfig.MysqlLog})
	}

	if articleDetailReq.ArticleID == "" {
		return nil, errors.New("")
	}

	var articleDetail common.Article
	err := DB.Model(&common.Article{}).Where("id = ?", articleDetailReq.ArticleID).Find(&articleDetail).Error
	if err != nil {
		return nil, err
	}

	return &articleDetail, nil
}
