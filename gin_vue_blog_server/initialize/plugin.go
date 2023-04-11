package initialize

import (
	"fmt"
	"gin_vue_blog_server/middleware"
	"github.com/gin-gonic/gin"
)

func InstallPlugin(r *gin.Engine) {
	//鉴权 JWT 配置
	PublicGroup := r.Group("")
	fmt.Println("无鉴权路由==》", PublicGroup)
	PrivateGroup := r.Group("")
	fmt.Println("有鉴权路由==》", PrivateGroup)
	//20230411 进度
	PrivateGroup.Use(middleware.JwtAuth())

}
