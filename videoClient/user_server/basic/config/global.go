package config

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config AppConfig
	Log    = zap.NewExample()
	Rdb    *redis.Client
)
