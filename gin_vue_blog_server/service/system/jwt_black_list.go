package system

import (
	"context"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/system"
	"gin_vue_blog_server/utils"
	"go.uber.org/zap"
)

type JwtService struct{}

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlackList) (err error) {
	err = global.DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.Redis.Get(context.Background(), userName).Result()
	return redisJWT, err
}

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.Config.JWT.Expire)
	if err != nil {
		return err
	}
	timer := dr
	err = global.Redis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func (jwtService *JwtService) DataInBlacklist(jwt string) bool {
	//判断jwt是否在黑名单内部
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

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
