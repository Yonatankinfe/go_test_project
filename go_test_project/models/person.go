package models

import (
    "github.com/google/uuid"
)

type Person struct {
    ID      string   `json:"id"`
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Hobbies []string `json:"hobbies"`
}

func NewPerson(name string, age int, hobbies []string) Person {
    return Person{
        ID:      uuid.New().String(),
        Name:    name,
        Age:     age,
        Hobbies: hobbies,
    }
}
