package main

import (
    "log"
    "github.com/Yonatankinfe/go_test_project/config/config.go"
    "github.com/Yonatankinfe/go_test_project/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.Use(config.CORS())
    router.Use(config.Handle500())
    personStorage := config.NewInMemoryPersonStorage()
    routes.SetupPersonRoutes(router, personStorage)

    router.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{"error": "Resource not found"})
    })

    if err := router.Run(config.ServerAddress()); err != nil {
        log.Fatalf("Failed to start the server: %v", err)
    }
}
