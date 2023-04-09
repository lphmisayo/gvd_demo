package global

import (
	"gin_vue_blog_demo/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Logger *zap.Logger
)
