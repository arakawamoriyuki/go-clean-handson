package entiry

import (
	vo "main/domain/value-object"
)

type User struct {
	Id   vo.UserId
	Name vo.UserName
}

func NewUser(id vo.UserId, name vo.UserName) *User {
	user := &User{
		Id:   id,
		Name: name,
	}
	return user
}

func (u User) Equals(user *User) bool {
	return u.Id == user.Id
}

func (u *User) ChangeName(name vo.UserName) {
	u.Name = name
}
