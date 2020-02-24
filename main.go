package main

import (
	service "github.com/Sankara98/gogo-service/service"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	server := service.NewServer()
	server.Run(":" + port)
}
