package repository

import (
	"fmt"
	"log"
	"main/domain/application/repository"
	"main/domain/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseInterface interface {
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}

// Data Access
type TodoRepository struct {
	db DatabaseInterface
}

func NewTodoRepository() repository.TodoRepositoryInterface {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"root",
		"pass",
		"localhost",
		"3306",
		"ca-sample",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Get(id int) (*entity.Todo, error) {

	var todo *entity.Todo
	if result := r.db.First(&todo, id); result.Error != nil {
		return nil, result.Error
	}

	return todo, nil
}
