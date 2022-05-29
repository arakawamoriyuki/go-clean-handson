package domainservice

import (
	entity "main/domain/entity"
	repository "main/domain/repository"
)

func UserExists(user entity.User) bool {
	userRepository := repository.NewUserRepository()

	exists := userRepository.Exists(user)

	return exists
}
