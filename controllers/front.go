package controllers

import (
	"encoding/json"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()
	tc := newTodoController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
	http.Handle("/todos", *tc)
	http.Handle("/todos/", *tc)
}

func encodeResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
