package main

import (
	"os"

	service "github.com/Sankara98/gogo-service/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	server := service.NewServer()
	server.Run(":" + port)
}
