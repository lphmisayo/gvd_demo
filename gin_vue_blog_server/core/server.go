package core

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/initialize"
	"gin_vue_blog_server/service/system"
)

func RunBackGroundServer() {
	if global.Config.Server.UseMultipoint || global.Config.Redis.Enable {
		//初始化Redis服务
		//initialize.Redis()
	}

	//从DB加载JWT数据
	if global.DB != nil {
		system.LoadAll()
	}

	router := initialize.Routers()
	router.Run(":8081")
}
