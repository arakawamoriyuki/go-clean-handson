package usecase

import (
	"main/domain/application/repository"
)

// Use Case Interactor
type TodoInteractor struct {
	OutputPort TodoOutputPortInterface
	Repository repository.TodoRepositoryInterface
}

func NewTodoInteractor(
	outputPort TodoOutputPortInterface,
	repository repository.TodoRepositoryInterface,
) TodoInputPortInterface {
	return &TodoInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (i *TodoInteractor) Get(inputData TodoInputData) (*TodoResponse, error) {

	todo, err := i.Repository.Get(inputData.Id)
	if err != nil {
		return nil, err
	}

	outputData := &TodoOutputData{
		Id:   todo.Id,
		Name: todo.Name,
	}

	res, err := i.OutputPort.Convert(*outputData)
	if err != nil {
		return nil, err
	}

	return res, nil
}
