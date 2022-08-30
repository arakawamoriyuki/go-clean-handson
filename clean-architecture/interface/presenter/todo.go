package presenter

import (
	"main/domain/application/usecase"
)

// Presenter
type TodoPresenter struct {
}

func NewTodoPresenter() usecase.TodoOutputPortInterface {
	return &TodoPresenter{}
}

func (p *TodoPresenter) Convert(outputData usecase.TodoOutputData) (*usecase.TodoResponse, error) {
	return &usecase.TodoResponse{
		Id:   outputData.Id,
		Name: outputData.Name,
	}, nil
}
