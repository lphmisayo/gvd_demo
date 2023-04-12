package system

import "gin_vue_blog_server/service"

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
)
