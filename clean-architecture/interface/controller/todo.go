package controller

import (
	"main/domain/application/usecase"
	"net/http"
	"strconv"
)

// Controller
type TodoController struct {
	InputPort usecase.TodoInputPortInterface
}

func NewTodoController(inputPort usecase.TodoInputPortInterface) *TodoController {
	return &TodoController{
		InputPort: inputPort,
	}
}

func (ctrl *TodoController) GetTodo(context Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		context.Abort()
		return
	}

	inputData := &usecase.TodoInputData{
		Id: id,
	}

	res, err := ctrl.InputPort.Get(*inputData)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, res)
}
