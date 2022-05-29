package main

import (
	"fmt"
	application "main/domain/application"
	repository "main/domain/repository"
)

func main() {
	repo := repository.NewUserRepository()
	app := application.NewUserApplication(repo)

	user, err := app.FindUser("Moriyuki Arakawa")
	if err != nil {
		fmt.Println(err)
		return
	}

	log := fmt.Sprintf("FindUser id:%d name:%s", user.Id, user.Name)
	fmt.Println(log)
}
