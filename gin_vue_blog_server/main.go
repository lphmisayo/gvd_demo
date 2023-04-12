package main

import (
	"gin_vue_blog_server/core"
	"gin_vue_blog_server/initialize"
)

func main() {
	ok := initialize.InitGlobalVariable() //初始化服务配置
	if ok {
		return
	}
	core.RunBackGroundServer()
}
