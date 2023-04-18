package system

import (
	v1 "gin_vue_blog_server/api/v1"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

func (*ArticleRouter) InitArticleRouter(group *gin.RouterGroup) (ri gin.IRouter) {
	articleRouter := group.Group("/article")

	articleListApi := v1.ApiGroupApp.ArticleApiGroup.ArticleListApi
	articleSingleApi := v1.ApiGroupApp.ArticleApiGroup.ArticleSingleApi

	{
		articleRouter.GET("/auth_articles", articleListApi.ArticleUserListView)
		articleRouter.GET("/detail", articleSingleApi.GetArticleDetailInfo)
	}
	return articleRouter
}
