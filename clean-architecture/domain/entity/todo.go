package entity

// Entities
type TodoInterface interface {
}

func NewTodo() TodoInterface {
	return &Todo{}
}

type Todo struct {
	Id   int
	Name string
}

type TodoForm struct {
	Name string
}
