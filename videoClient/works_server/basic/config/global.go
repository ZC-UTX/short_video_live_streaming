package config

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config AppConfig
	Log    = zap.NewExample()
)
