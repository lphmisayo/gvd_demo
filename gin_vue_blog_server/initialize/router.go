package initialize

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/middleware"
	"gin_vue_blog_server/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.Default()
	//InstallPlugin(Router) // 安装 插件

	systemRouterGroup := routers.RouterGroupApp.System

	//头像文件设置静态地址

	//swagger

	//配置路由
	PublicGroup := r.Group(global.Config.System.RouterPrefix)
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "连接正常")
		})
	}
	{
		systemRouterGroup.InitBaseRouter(PublicGroup) //初始化基础路由
		systemRouterGroup.InitFileRouter(PublicGroup) //初始化文件路由
	}
	PrivateGroup := r.Group(global.Config.System.RouterPrefix)
	PrivateGroup.Use(middleware.JwtAuth()).Use(middleware.CasbinRabc())
	return r
}
