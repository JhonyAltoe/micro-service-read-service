package main

import (
	"github.com/JhonyAltoe/micro-service-read-service/handlers"
	"github.com/JhonyAltoe/micro-service-read-service/logs"
	"github.com/sirupsen/logrus"
)

var file = logs.Init()

func main() {
	defer file.Close()
	handlers.Repo.Ping()
	r := handlers.Router()
	r.GET("/services/:company", handlers.GetServices)
	if err := r.Run(":8080"); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}