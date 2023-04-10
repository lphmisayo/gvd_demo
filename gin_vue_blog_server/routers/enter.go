package routers

import (
	"fmt"
	"gin_vue_blog_server/core"
	"gin_vue_blog_server/flag"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/initialize"
	"gin_vue_blog_server/utils"
)

// InitGlobalVariable 初始化全局变量
func InitGlobalVariable() {
	utils.InitConfig() //初始化配置文件
	fmt.Println(global.Config)
	core.InitLogger() //初始化日志文件
	global.DB = initialize.InitMysqlDB()
	global.Viper = core.Viper()

	option := flag.Parse()
	if flag.IsWebStop(option) { //需要执行后重启服务的指令，将进入此方法
		flag.SwitchOption(option)
		return
	}
}

/*func BackServer() *http.Server {
	backPort := global.Config.Server.BackPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      BackRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
*/
