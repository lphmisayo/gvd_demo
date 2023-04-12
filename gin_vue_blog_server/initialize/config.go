package initialize

import (
	"gin_vue_blog_server/flag"
	"gin_vue_blog_server/global"
)

// InitGlobalVariable 初始化全局变量
func InitGlobalVariable() (stopServer bool) {
	//初始化命令行操作
	global.Option = flag.Parse()

	//初始化配置文件 Viper
	global.Viper = Viper(global.Option.Config)

	//初始化数据库
	global.DB = InitMysqlDB()
	//fmt.Println(global.Config)

	//初始化日志文件
	InitLogger()
	stopServer = flag.IsWebStop(global.Option)
	if stopServer { //需要执行后重启服务的指令，将进入此方法
		flag.SwitchOption(global.Option)
	}

	return stopServer
}
