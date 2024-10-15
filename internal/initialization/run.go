package initialization

import (
	"backend-go/global"
	"fmt"
)

func Run() {
	// load configuration
	LoadConfig()
	mysql := global.Config.Mysql
	fmt.Println("Loading configuration mysql", mysql.User, mysql.Password)
		
	InitMysql()
	InitRedis()
	InitLogger()

	r := InitRouter()
	r.Run(":8080")
}