package core

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/initialize"
	"gin_vue_blog_server/service/system"
)

func RunServer() {
	if global.Config.Server.UseMultipoint || global.Config.Redis.Enable {
		initialize.Redis()
	}

	if global.DB != nil {
		system.LoadAll()
	}
}
