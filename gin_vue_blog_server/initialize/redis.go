package initialize

import (
	"context"
	"gin_vue_blog_server/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("Redis Connect Ping Failed, 错误原因：", zap.Error(err))
	} else {
		global.Logger.Info("Redis Connect Ping Response：", zap.String("ping", ping))
		global.Redis = client
	}
}
