package global

import (
	"backend-go/pkg/logger"
	"backend-go/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	MysqlDB *gorm.DB
)