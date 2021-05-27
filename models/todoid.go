package models

type TodoId struct {
	Value string
}

func (t TodoId) ModelName() string { return "todoid" }
