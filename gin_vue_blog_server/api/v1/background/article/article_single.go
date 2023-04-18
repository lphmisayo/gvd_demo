package article

import (
	"gin_vue_blog_server/model/request"
	"gin_vue_blog_server/model/response"
	"github.com/gin-gonic/gin"
)

func (*ArticleSingleApi) GetArticleDetailInfo(c *gin.Context) {
	var articleDetailReq request.ArticleDetailReq
	err := c.ShouldBindJSON(&articleDetailReq)
	if err != nil {
		response.FailWithCode(response.ParamError, c)
		return
	}

	articleDetail, err := articleSingleService.GetArticleDetail(articleDetailReq, true)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(articleDetail, c)
}
