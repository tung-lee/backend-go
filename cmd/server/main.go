package main

import "backend-go/internal/routers"

func main() {
	r := routers.NewRouter()

	r.Run(":8082") // default listen and serve on 0.0.0.0:8080
}

