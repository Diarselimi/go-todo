package controllers

import (
	"net/http"
	"regexp"

	"github.com/Diarselimi/go-todo/models"
	"github.com/Diarselimi/go-todo/requestparser"
)

type todoController struct {
	todoIDPattern *regexp.Regexp
}

func (tc todoController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		saveOne(w, r)
	} else {
		getAll(w)
	}
}

func saveOne(w http.ResponseWriter, r *http.Request) {
	todo, err := requestparser.ParseTodo(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	models.AddTodo(&todo)
	encodeResponse(todo, w)
}

func getAll(w http.ResponseWriter) {
	encodeResponse(models.GetAll(), w)
}

func newTodoController() *todoController {
	return &todoController{
		todoIDPattern: regexp.MustCompile(`^/todos/(\d+)/?`),
	}
}
