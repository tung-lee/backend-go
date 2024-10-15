package global

import (
	"backend-go/pkg/logger"
	"backend-go/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)