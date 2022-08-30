package usecase

// Output Data <DS>
type TodoOutputData struct {
	Id   int
	Name string
}

type TodoResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Output Boundary <I>
type TodoOutputPortInterface interface {
	Convert(TodoOutputData) (*TodoResponse, error)
}
