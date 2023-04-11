package initialize

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	router := gin.Default()
	//InstallPlugin(Router) // 安装 插件
	return router
}
