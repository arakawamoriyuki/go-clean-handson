package api

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetTodoRequest struct {
	ID int `uri:"id" json:"id" binding:"required"`
}

type CreateTodoRequest struct {
	Name string `form:"name" binding:"required"`
	Done *bool  `form:"done" binding:"required"`
}

type UpdateTodoRequest struct {
	ID   int     `uri:"id" binding:"required"`
	Name *string `form:"name"`
	Done *bool   `form:"done"`
}

type DeleteTodoRequest struct {
	ID int `uri:"id" json:"id" binding:"required"`
}

func GetTodos(c *gin.Context) {
	todos, err := models.GetTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	req := GetTodoRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := models.GetTodo(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	req := CreateTodoRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := models.CreateTodo(&models.Todo{
		Name: req.Name,
		Done: *req.Done,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	req := UpdateTodoRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := models.UpdateTodo(&models.Todo{
		ID:   req.ID,
		Name: *req.Name,
		Done: *req.Done,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	req := DeleteTodoRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DeleteTodo(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.Status(http.StatusNoContent)
}
