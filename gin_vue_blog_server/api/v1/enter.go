package v1

import (
	"gin_vue_blog_server/api/v1/background/article"
	"gin_vue_blog_server/api/v1/background/file"
	"gin_vue_blog_server/api/v1/background/system"
	"gin_vue_blog_server/api/v1/background/user"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	FileApiGroup    file.ApiGroup
	ArticleApiGroup article.ApiGroup
	UserApiGroup    user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
