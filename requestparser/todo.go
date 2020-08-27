package requestparser

import (
	"encoding/json"
	"net/http"

	"github.com/Diarselimi/go-todo/models"
)

func ParseTodo(r *http.Request) (models.Todo, error) {
	decoder := json.NewDecoder(r.Body) //get the decoder
	var todo models.Todo               // declare todo

	err := decoder.Decode(&todo) //decode
	if err != nil {
		return models.Todo{}, err //throw error if something is wrong
	}

	return todo, nil // return decoded
}
