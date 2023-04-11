package utils

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/system/request"
	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*request.Claims, error) {
	token := c.Request.Header.Get("x-token")
	jwt := NewJwt()
	claims, err := jwt.ParseToken(token)
	if err != nil {
		global.Logger.Error("获取Jwt解析信息失败，请检查请求头是否存在x-token且claims是否为规范结构")
	}
	return claims, err
}
