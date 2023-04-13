package core

import (
	"fmt"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/initialize"
	"gin_vue_blog_server/service/system"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsBackGroundServer() {
	if global.Config.Server.UseMultipoint || global.Config.Redis.Enable {
		//初始化Redis服务
		//initialize.Redis()
	}

	//从DB加载JWT数据
	if global.DB != nil {
		system.LoadAll()
	}

	router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	server := initWinServer(address, router)
	//保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.Logger.Info("服务启动成功！ ", zap.String("address", address))
	//router.Run(":8081")
	//启动服务
	global.Logger.Error(server.ListenAndServe().Error())
}
