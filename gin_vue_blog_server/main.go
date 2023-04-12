package main

import (
	"gin_vue_blog_server/core"
	"gin_vue_blog_server/initialize"
)

func main() {
	initialize.InitGlobalVariable() //初始化服务配置
	core.RunBackGroundServer()
}
