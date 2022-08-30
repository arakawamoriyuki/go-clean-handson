package api

import (
	"main/interface/controller"

	"github.com/gin-gonic/gin"
)

func GetTodoHandler(ctrl *controller.TodoController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctrl.GetTodo(c)
	}
}
