package inte

import (
	"gin_vue_blog_server/model/response"
	"github.com/gin-gonic/gin"
)

func (b *BaseApi) Login(c *gin.Context) {
	response.OK(c)
}
