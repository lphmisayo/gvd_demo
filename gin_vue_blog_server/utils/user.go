package utils

import (
	"fmt"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/common"
)

func CheckIsExist(username string) (bool, error) {
	if nil == global.DB {
		return true, fmt.Errorf("数据库未初始化")
	}
	var user common.User
	err := global.DB.Model(common.User{}).Select("username = ?", username).Scan(&user).Error
	if err != nil {
		return true, err
	}
	if user.UserName != "" && user.UserName == username {
		return true, fmt.Errorf("用户名已存在！")
	}
	return false, nil
}
