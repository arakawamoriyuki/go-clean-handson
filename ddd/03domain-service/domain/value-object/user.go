package valueobject

import (
	"errors"
)

type UserId struct {
	value int
}

func NewUserId(id int) (*UserId, error) {
	if id < 1 {
		return nil, errors.New("ユーザーIDは1以上である必要があります")
	}

	userId := &UserId{
		value: id,
	}
	return userId, nil
}

type UserName struct {
	value string
}

func NewUserName(name string) (*UserName, error) {
	if len(name) < 3 {
		return nil, errors.New("ユーザー名は3文字以上である必要があります")
	}

	userName := &UserName{
		value: name,
	}
	return userName, nil
}
