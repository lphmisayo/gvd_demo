package service

import (
	"gin_vue_blog_server/service/file"
	"gin_vue_blog_server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	FileServiceGroup   file.ServiceGroup
}
