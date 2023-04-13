package v1

import (
	"gin_vue_blog_server/api/v1/background/file"
	"gin_vue_blog_server/api/v1/background/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	FileApiGroup   file.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
