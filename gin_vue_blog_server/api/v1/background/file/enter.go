package file

import "gin_vue_blog_server/service"

type FileApi struct {
}

type ApiGroup struct {
	FileApi
}

var (
	localFileService = service.ServiceGroupApp.FileServiceGroup.LocalFileService
)
