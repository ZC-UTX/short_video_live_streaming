package init

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"user_server/basic/config"
)

func init() {
	InitConfig()
	InitMysql()
}

func InitConfig() {
	viper.SetConfigFile("../config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		config.Log.Error("viper.ReadInConfig failed", zap.Error(err))
	}
	err = viper.Unmarshal(&config.Config)
	if err != nil {
		config.Log.Error("viper.Unmarshal failed", zap.Error(err))
	}
	config.Log.Info(fmt.Sprintf("Video Config %+v", config.Config))
}

func InitMysql() {
	var err error
	conf := config.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		config.Log.Error("gorm.Open failed", zap.Error(err))
	}
	config.Log.Info(fmt.Sprintf("Mysql connect success, %+v", config.DB))

	sqlDB, err := config.DB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

}
