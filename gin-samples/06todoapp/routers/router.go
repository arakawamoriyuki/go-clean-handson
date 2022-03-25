package routers

import (
	"github.com/gin-gonic/gin"

	"main/routers/api"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("api")
	{
		todosGroup := apiGroup.Group("todos")
		{
			todosGroup.GET("", api.GetTodos)
			todosGroup.GET(":id", api.GetTodo)
			todosGroup.POST("", api.CreateTodo)
			todosGroup.PATCH(":id", api.UpdateTodo)
			todosGroup.DELETE(":id", api.DeleteTodo)
		}
	}

	return router
}
