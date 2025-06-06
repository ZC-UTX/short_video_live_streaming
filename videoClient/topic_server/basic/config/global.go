package config

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Log   = zap.NewExample()
	Nacos string
	DB    *gorm.DB
	Conf  AppConfig
	RDB   *redis.Client
)
