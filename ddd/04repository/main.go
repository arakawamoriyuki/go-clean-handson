package main

import (
	"fmt"
	ds "main/domain/domain-service"
	entity "main/domain/entity"
	vo "main/domain/value-object"
)

func main() {
	userId, err := vo.NewUserId(1)
	if err != nil {
		fmt.Println(err)
	}
	userName, err := vo.NewUserName("新川 盛幸")
	if err != nil {
		fmt.Println(err)
	}
	user := entity.NewUser(*userId, *userName)

	log := fmt.Sprintf("Exists: %t", ds.UserExists(*user))
	fmt.Println(log)
}
