package system

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/system"
	"go.uber.org/zap"
)

func LoadAll() {
	var data []string
	err := global.DB.Model(&system.JwtBlackList{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.Logger.Error("加载数据库JWT黑名单失败！", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		// jwt黑名单 加入 BlackCache 中
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
