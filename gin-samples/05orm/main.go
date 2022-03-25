package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func main() {
	dsn := "root:pass@tcp(localhost:3306)/gin-sample"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create
	if result := db.Create(&Todo{Name: "test", Done: false}); result.Error != nil {
		log.Fatal(result.Error)
		return
	}

	// Read
	var todo Todo
	if result := db.First(&todo, 1); result.Error != nil {
		log.Fatal(result.Error)
		return
	}
	fmt.Println(todo.Name)

	if result := db.First(&todo, "name = ?", "test"); result.Error != nil {
		log.Fatal(result.Error)
		return
	}
	fmt.Println(todo.Name)

	// Update
	if result := db.Model(&todo).Update("name", "changed"); result.Error != nil {
		log.Fatal(result.Error)
		return
	}
	fmt.Println(todo.Name)

	if result := db.Model(&todo).Updates(&Todo{Name: "changed 2", Done: true}); result.Error != nil {
		log.Fatal(result.Error)
		return
	}
	fmt.Println(todo.Name)
	fmt.Println(todo.Done)

	if result := db.Model(&todo).Updates(map[string]interface{}{"name": "changed 3", "done": false}); result.Error != nil {
		log.Fatal(result.Error)
		return
	}
	fmt.Println(todo.Name)
	fmt.Println(todo.Done)

	// Delete (soft delete)
	if result := db.Delete(&todo, 1); result.Error != nil {
		log.Fatal(result.Error)
		return
	}
}
