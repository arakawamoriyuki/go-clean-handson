package router

import (
	"github.com/gin-gonic/gin"

	"main/domain/application/usecase"
	"main/infrastructure/router/api"
	"main/interface/controller"
	"main/interface/presenter"
	"main/interface/repository"
)

func SetupRouter() *gin.Engine {
	todoPresenter := presenter.NewTodoPresenter()
	todoRepository := repository.NewTodoRepository()
	todoInteractor := usecase.NewTodoInteractor(todoPresenter, todoRepository)
	todoController := controller.NewTodoController(todoInteractor)

	router := gin.Default()
	apiGroup := router.Group("api")
	{
		todosGroup := apiGroup.Group("todos")
		{
			// todosGroup.GET("", api.GetTodos)
			todosGroup.GET(":id", api.GetTodoHandler(todoController))
			// todosGroup.POST("", api.CreateTodo)
			// todosGroup.PATCH(":id", api.UpdateTodo)
			// todosGroup.DELETE(":id", api.DeleteTodo)
		}
	}

	return router
}
