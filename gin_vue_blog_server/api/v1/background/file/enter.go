package file

import "gin_vue_blog_server/service"

type ImageApi struct {
}

type ApiGroup struct {
	ImageApi
}

var (
	localFileService  = service.ServiceGroupApp.FileServiceGroup.LocalFileService
	localImageService = service.ServiceGroupApp.FileServiceGroup.LocalImageService
)
