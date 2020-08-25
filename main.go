package main

import (
	"fmt"

	"github.com/Diarselimi/go-todo/models"
)

func main() {
	me := models.User{
		ID:        1,
		FirstName: "Test",
		LastName:  "Selimi",
	}
	fmt.Println("Here we GO!", me)
}
