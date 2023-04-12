package global

import (
	"gin_vue_blog_server/config"
	"gin_vue_blog_server/config/config_model"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	Config             *config.Config
	DB                 *gorm.DB
	Logger             *zap.Logger
	Viper              *viper.Viper
	Redis              *redis.Client
	BlackCache         local_cache.Cache
	ConcurrencyControl = &singleflight.Group{}
	Option             *config_model.Option
)
