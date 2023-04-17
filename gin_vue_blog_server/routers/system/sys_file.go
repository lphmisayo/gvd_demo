package system

import (
	v1 "gin_vue_blog_server/api/v1"
	"github.com/gin-gonic/gin"
)

type FileRouter struct {
}

func (*FileRouter) InitFileRouter(Router *gin.RouterGroup) (ri gin.IRouter) {
	fileRouter := Router.Group("file")
	fileApi := v1.ApiGroupApp.FileApiGroup.ImageApi
	{
		fileRouter.GET("/images", fileApi.ImageListView)
		fileRouter.POST("/images", fileApi.ImageUploadView)
		fileRouter.PUT("/images", fileApi.ImageUpdateView)
		fileRouter.DELETE("/images", fileApi.ImageDeleteView)
	}
	return fileRouter
}
