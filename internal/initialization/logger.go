package initialization

import (
	"backend-go/global"
	"backend-go/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}