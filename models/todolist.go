package models

type TodoList struct {
	Todos *[]Todo
}

func (t TodoList) ModelName() string { return "todolist" }
