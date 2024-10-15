package initialization

import (
	"backend-go/global"
	"fmt"

	"go.uber.org/zap"
)

func Run() {
	// load configuration
	LoadConfig()
	mysql := global.Config.Mysql
	fmt.Println("Loading configuration mysql", mysql.User, mysql.Password)
		
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "success"))

	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8080")
}