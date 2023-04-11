package middleware

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/response"
	model_system "gin_vue_blog_server/model/system"
	"gin_vue_blog_server/service"
	"gin_vue_blog_server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDataAndMsg(gin.H{"reload": true}, "未登录或非法访问！", c)
			c.Abort() //中间件出现问题，用c.Abort终止服务
			return
		}
		if jwtService.DataInBlacklist(token) {
			response.FailWithDataAndMsg(gin.H{"reload": true}, "您的账户异地登陆或者令牌失效，请重新登陆！", c)
			c.Abort()
			return
		}
		jwt := utils.NewJwt()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == utils.ExpiredToken {
				response.FailWithDataAndMsg(gin.H{"reload": true}, "授权已过期！", c)
				c.Abort()
				return
			}
			response.FailWithDataAndMsg(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			duration, _ := utils.ParseDuration(global.Config.JWT.Expire)
			claims.ExpiresAt = time.Now().Add(duration).Unix()
			newToken, _ := jwt.CreateTokenByOldToken(token, *claims)
			newClaims, _ := jwt.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.Config.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.Logger.Error("获取Redis配置失败", zap.Error(err))
				} else {
					_ = jwtService.JsonInBlacklist(model_system.JwtBlackList{Jwt: RedisJwtToken})
				}
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
