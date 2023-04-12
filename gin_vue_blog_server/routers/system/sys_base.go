package system

import (
	v1 "gin_vue_blog_server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (ri gin.IRouter) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		//baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("login", baseApi.LoginDefault)
		baseRouter.POST("register", baseApi.Register)
	}
	return baseRouter
}
