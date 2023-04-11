package initialize

import (
	"fmt"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/middleware"
	"gin_vue_blog_server/plugin/email"
	"gin_vue_blog_server/utils"
	"github.com/gin-gonic/gin"
)

func InstallPlugin(r *gin.Engine) {
	//鉴权 JWT 配置
	PublicGroup := r.Group("")
	fmt.Println("无鉴权路由==》", PublicGroup)
	PrivateGroup := r.Group("")
	fmt.Println("有鉴权路由==》", PrivateGroup)
	PrivateGroup.Use(middleware.JwtAuth()).Use(middleware.CasbinRabc())
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.Config.Email.To,
		global.Config.Email.From,
		global.Config.Email.Host,
		global.Config.Email.Secret,
		global.Config.Email.Nickname,
		global.Config.Email.Port,
		global.Config.Email.IsSSL,
	))
}

func PluginInit(group *gin.RouterGroup, Plugin ...utils.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}
