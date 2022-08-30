package repository

import (
	entiry "main/domain/entity"
)

// Data Access Interface <I>
type TodoRepositoryInterface interface {
	// GetList() ([]*entiry.Todo, error)
	Get(int) (*entiry.Todo, error)
	// Create(*entiry.TodoForm) (*entiry.Todo, error)
	// Update(*entiry.Todo) (*entiry.Todo, error)
	// Delete(entiry.TodoId) error
}
