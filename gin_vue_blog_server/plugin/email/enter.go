package email

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/plugin/email/router"
	"github.com/gin-gonic/gin"
)

type emailPlugin struct{}

func CreateEmailPlug(To, From, Host, Secret, Nickname string, Port int, IsSSL bool) *emailPlugin {
	global.Config.Email.To = To
	global.Config.Email.From = From
	global.Config.Email.Host = Host
	global.Config.Email.Secret = Secret
	global.Config.Email.Nickname = Nickname
	global.Config.Email.Port = Port
	global.Config.Email.IsSSL = IsSSL
	return &emailPlugin{}
}

func (*emailPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitEmailRouter(group)
}

func (*emailPlugin) RouterPath() string {
	return "email"
}
