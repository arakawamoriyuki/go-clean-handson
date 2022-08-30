package usecase

// Input Data <DS>
type TodoInputData struct {
	Id int
}

// Input Boundary <I>
type TodoInputPortInterface interface {
	Get(TodoInputData) (*TodoResponse, error)
}
