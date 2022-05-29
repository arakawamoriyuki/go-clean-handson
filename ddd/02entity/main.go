package main

import (
	"fmt"
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

	// 可変である
	//   ので名前を変更することができる
	newUserName, err := vo.NewUserName("Moriyuki Arakawa")
	if err != nil {
		fmt.Println(err)
	}
	log := fmt.Sprintf("旧User id:%d name:%s", user.Id, user.Name)
	fmt.Println(log)
	user.ChangeName(*newUserName)
	log = fmt.Sprintf("新User id:%d name:%s", user.Id, user.Name)
	fmt.Println(log)

	// 同じ属性であっても区別される
	//   のでIDが同じではない限り区別される
	userId2, err := vo.NewUserId(2)
	if err != nil {
		fmt.Println(err)
	}
	userName2, err := vo.NewUserName("Moriyuki Arakawa")
	if err != nil {
		fmt.Println(err)
	}
	user2 := entity.NewUser(*userId2, *userName2)
	log = fmt.Sprintf("Equals: %t", user.Equals(user2))
	fmt.Println(log)

	// 同一性により区別される
	//   のでIDが同じなら同一とされる
	user3 := entity.NewUser(*userId, *userName2)
	log = fmt.Sprintf("Equals: %t", user.Equals(user3))
	fmt.Println(log)
}
