package middleware

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/response"
	"gin_vue_blog_server/service"
	"gin_vue_blog_server/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

func CasbinRabc() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.Config.System.Env != "develop" {
			waitUse, _ := utils.GetClaims(c)
			// 获取请求的PATH
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, global.Config.System.RouterPrefix)
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			sub := strconv.Itoa(int(waitUse.AuthorityId))
			e := casbinService.Casbin()
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDataAndMsg(gin.H{}, "权限不足！", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
