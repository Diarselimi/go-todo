package models

import (
	"fmt"
)

type Todo struct {
	ID     int
	Name   string
	IsDone bool
}

var (
	todos      []*Todo
	nextTodoID = 1
)

func AddTodo(todo *Todo) error {
	for _, t := range todos {
		if todo.ID == t.ID {
			return fmt.Errorf("user with id '%v' already exists", t.ID)
		}
	}
	assignID(todo)
	todos = append(todos, todo)
	return nil
}

func GetAll() []*Todo {
	return todos
}

func assignID(todo *Todo) {
	todo.ID = nextTodoID
	nextTodoID++
}
