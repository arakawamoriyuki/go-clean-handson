package repository

import (
	entity "main/domain/entity"
	vo "main/domain/value-object"
)

type UserRepositoryInterface interface {
	Exists(user entity.User) bool
	Find(name string) (*entity.User, error)
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

func (r *UserRepository) Find(name string) (*entity.User, error) {

	// TODO: DBからuser.nameを検索してuserオブジェクトを生成して返す
	userId, err := vo.NewUserId(1)
	if err != nil {
		return nil, err
	}
	userName, err := vo.NewUserName("Moriyuki Arakawa")
	if err != nil {
		return nil, err
	}
	user := entity.NewUser(*userId, *userName)

	return user, nil
}
