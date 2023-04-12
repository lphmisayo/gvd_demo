package system

import (
	"fmt"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/common"
	"gin_vue_blog_server/model/system"
	sys_res "gin_vue_blog_server/model/system/response"
	"gin_vue_blog_server/utils"
	"time"
)

type UserService struct{}

func (UserService) Register(r *common.User) error {
	if nil == global.DB {
		return fmt.Errorf("数据库未初始化")
	}
	isExist, err := utils.CheckIsExist(r.UserName)
	if isExist {
		return err
	}
	r.CreateAt = time.Now()
	r.UpdateAt = time.Now()
	err = global.DB.Create(r).Error
	if err != nil {
		return err
	}
	return nil
}

func (UserService) Login(u *system.SysUser) (user *system.SysUser, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("数据库未初始化")
	}
	//var userIne system.SysUser
	err = global.DB.Where("username = ?", u.Username).
		Preload("Authorities").
		Preload("Authority").
		First(user).Error
	if err == nil {
		if ok := utils.CheckBcrypt(u.Password, user.Password); !ok {
			return nil, fmt.Errorf("密码错误！")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(user)
	}
	return user, err
}

func (UserService) LoginDefault(u *common.User) (user *sys_res.User, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("数据库未初始化")
	}
	err = global.DB.Where("username = ? and password = ?", u.UserName, u.PassWord).First(&user).Error
	user.CreateTime = user.CreateAt.Format("2006-01-02 15:04:05")
	return user, err
}
