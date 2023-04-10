package global

import (
	"gin_vue_blog_server/config"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config     *config.Config
	DB         *gorm.DB
	Logger     *zap.Logger
	Viper      *viper.Viper
	Redis      *redis.Client
	BlackCache local_cache.Cache
)
