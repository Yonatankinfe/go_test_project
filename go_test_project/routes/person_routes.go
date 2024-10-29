package routes

import (
    "github.com/Ephrem-shimels21/GoCrudChallenge/handlers"
    "github.com/Ephrem-shimels21/GoCrudChallenge/storage"
    "github.com/gin-gonic/gin"
)

func SetupPersonRoutes(router *gin.Engine, storage *storage.InMemoryPersonStorage) {
    router.GET("/person", handlers.GetPersons(storage))
    router.GET("/person/:id", handlers.GetPerson(storage))
    router.POST("/person", handlers.AddPerson(storage))
    router.PUT("/person/:id", handlers.UpdatePerson(storage))
    router.DELETE("/person/:id", handlers.DeletePerson(storage))
}