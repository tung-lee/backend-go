package global

import (
	"backend-go/pkg/logger"
	"backend-go/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	MysqlDB *gorm.DB
	RedisDB *redis.Client
)