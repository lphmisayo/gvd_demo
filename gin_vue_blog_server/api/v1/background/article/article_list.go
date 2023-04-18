package article

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model"
	"gin_vue_blog_server/model/common"
	"gin_vue_blog_server/model/response"
	"gin_vue_blog_server/service/article"
	"github.com/gin-gonic/gin"
)

// ArticleUserListView 分页查询用户文章
func (*ArticleListApi) ArticleUserListView(c *gin.Context) {
	var pageInfo model.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithCode(response.ParamError, c)
		return
	}

	articleList, count, err := articleListService.GetAuthorArticlesSimpleInfo(common.Article{}, article.Option{
		PageInfo: pageInfo,
		Debug:    global.MysqlLogConfig.Debug,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithList(articleList, count, "文章查询成功！", c)
}
