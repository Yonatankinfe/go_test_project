package handlers

import (
    "github.com/Yonatankinfe/go_test_progect/models"
    "github.com/Yonatankinfe/go_test_project/storage"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetPersons(storage *storage.InMemoryPersonStorage) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        persons := storage.GetPersons()
        ctx.JSON(http.StatusOK, persons)
    }
}

func GetPerson(storage *storage.InMemoryPersonStorage) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        id := ctx.Param("id")
        person, err := storage.GetPersonByID(id)
        if err != nil {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
            return
        }
        ctx.JSON(http.StatusOK, person)
    }
}

func AddPerson(storage *storage.InMemoryPersonStorage) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var newPerson models.Person
        if err := ctx.BindJSON(&newPerson); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
            return
        }
        newPerson.ID = uuid.New().String()
        storage.AddPerson(newPerson)
        ctx.JSON(http.StatusCreated, newPerson)
    }
}

func UpdatePerson(storage *storage.InMemoryPersonStorage) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        id := ctx.Param("id")
        existingPerson, err := storage.GetPersonByID(id)
        if err != nil {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
            return
        }
        var updatedFields models.Person
        if err := ctx.BindJSON(&updatedFields); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
            return
        }

        if updatedFields.Name != "" {
            existingPerson.Name = updatedFields.Name
        }
        if updatedFields.Age != 0 {
            existingPerson.Age = updatedFields.Age
        }
        if updatedFields.Hobbies != nil {
            existingPerson.Hobbies = updatedFields.Hobbies
        }

        if err := storage.UpdatePerson(id, *existingPerson); err != nil {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
            return
        }
        ctx.JSON(http.StatusOK, gin.H{"message": "Person updated successfully", "person": existingPerson})
    }
}

func DeletePerson(storage *storage.InMemoryPersonStorage) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        id := ctx.Param("id")
        if err := storage.DeletePerson(id); err != nil {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
            return
        }
        ctx.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
    }
}
