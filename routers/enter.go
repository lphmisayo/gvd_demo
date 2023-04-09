package routers

import (
	"fmt"
	"gin_vue_blog_demo/dao"
	"gin_vue_blog_demo/global"
	"gin_vue_blog_demo/utils"
)

// InitGlobalVariable 初始化全局变量
func InitGlobalVariable() {
	utils.InitConfig() //初始化配置文件
	fmt.Println(global.Config)
	utils.InitLogger() //初始化日志文件
	global.DB = dao.InitMysqlDB()
}
