package application

import (
	entity "main/domain/entity"
	repository "main/domain/repository"
)

type UserApplication struct {
	repository repository.UserRepositoryInterface
}

func NewUserApplication(repository repository.UserRepositoryInterface) *UserApplication {
	userApplication := &UserApplication{
		repository: repository,
	}
	return userApplication
}

func (ua UserApplication) FindUser(name string) (*entity.User, error) {
	user, err := ua.repository.Find(name)

	if err != nil {
		return nil, err
	}

	return user, err
}
