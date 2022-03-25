package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int            `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	Done      bool           `json:"done"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func GetTodos() ([]*Todo, error) {
	var todos []*Todo
	if result := db.Find(&todos); result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func GetTodo(id int) (*Todo, error) {
	var todo *Todo
	if result := db.First(&todo, id); result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}

func CreateTodo(todo *Todo) (*Todo, error) {
	if result := db.Create(&todo); result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}

func UpdateTodo(todo *Todo) (*Todo, error) {
	if result := db.Model(&todo).Updates(map[string]interface{}{"name": todo.Name, "done": todo.Done}); result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}

func DeleteTodo(id int) error {
	// soft delete
	var todo *Todo
	if result := db.Delete(&todo, id); result.Error != nil {
		return result.Error
	}

	return nil
}
