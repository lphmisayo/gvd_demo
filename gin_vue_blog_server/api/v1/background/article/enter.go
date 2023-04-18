package article

import (
	"gin_vue_blog_server/service"
)

type ApiGroup struct {
	ArticleListApi
	ArticleSingleApi
}

type ArticleListApi struct{}
type ArticleSingleApi struct{}

var (
	articleListService   = service.ServiceGroupApp.ArticleServiceGroup.ArticleListService
	articleSingleService = service.ServiceGroupApp.ArticleServiceGroup.ArticleSingleService
)
