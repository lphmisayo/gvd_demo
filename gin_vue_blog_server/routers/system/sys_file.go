package system

import (
	v1 "gin_vue_blog_server/api/v1"
	"github.com/gin-gonic/gin"
)

type FileRouter struct {
}

func (FileRouter) InitFileRouter(Router *gin.RouterGroup) (ri gin.IRouter) {
	fileRouter := Router.Group("file")
	fileApi := v1.ApiGroupApp.FileApiGroup.FileApi
	{
		fileRouter.GET("/images", fileApi.ImageListView)
	}
	return fileRouter
}
