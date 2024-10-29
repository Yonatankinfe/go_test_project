package config

import (
    "github.com/Yonatankinfe/go_test_project/models/person.go"
    "github.com/Yonatankinfe/go_test_project/storage/person_storage.go"
    
    "github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
    return cors.Default()
}

func Handle500() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if r := recover(); r != nil {
                c.JSON(500, gin.H{"error": "Internal server error. Please try again later."})
                c.Abort()
            }
        }()
        c.Next()
    }
}

func NewInMemoryPersonStorage() *storage.InMemoryPersonStorage {
    return storage.NewInMemoryPersonStorage()
}

func ServerAddress() string {
    return ":8080"
}
