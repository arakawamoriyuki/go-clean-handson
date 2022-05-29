package repository

import (
	entity "main/domain/entity"
)

type UserRepositoryInterface interface {
	Exists(user entity.User) bool
}

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	repository := &UserRepository{}
	return repository
}

func (r *UserRepository) Exists(user entity.User) bool {

	// TODO: DBからuser.Idのレコードがあるか確認する処理
	exists := false

	return exists
}
