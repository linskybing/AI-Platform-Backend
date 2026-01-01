package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/linskybing/platform-go/internal/api/middleware"
	"github.com/linskybing/platform-go/internal/api/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggingMiddleware())

	routes.RegisterRoutes(router)

	port := ":8080"
	log.Printf("Starting API server on %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
