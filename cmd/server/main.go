package main

import "backend-go/internal/initialization"

func main() {
	// r := routers.NewRouter()

	// InitMysql()
	// InitRedis()
	// InitKafka()
	//...

	// r.Run(":8082") // default listen and serve on 0.0.0.0:8080

	initialization.Run()
}
