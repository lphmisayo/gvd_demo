package system

import (
	"errors"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/system"
	"gorm.io/gorm"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (menuService MenuService) UserAuthorityDefaultRouter(user *system.SysUser) {
	var menuIds []string
	err := global.DB.Model(&system.SysAuthorityMenu{}).
		Where("sys_authority_authority_id = ?", user.AuthorityId).
		Pluck("sys_base_menu_id", &menuIds).Error
	if err != nil {
		return
	}
	var menu system.SysBaseMenu
	err = global.DB.First(&menu, "name = ? and id in (?)", user.Authority.DefaultRouter, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Authority.DefaultRouter = "404"
	}
}
