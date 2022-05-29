package main

import (
	"fmt"
	vo "main/domain/value-object"
)

func main() {
	fullName1 := vo.NewFullName("Moriyuki", "Arakawa")

	// 不変である
	//   不変なので内部の値を変更するメソッドを持たない
	// fullName1.ChangeFirstName("盛幸")
	// fullName1.ChangeLastName("新川")

	// 交換可能である
	//   なので代入は可能
	fullName1 = vo.NewFullName("Moriyuki", "Arakawa")

	// 等価性によって評価される
	//   なので比較用メソッドを持っている
	fullName2 := vo.NewFullName("盛幸", "新川")
	fullName1.Equals(fullName2)
	log := fmt.Sprintf("Equals: %t", fullName1.Equals(fullName2))
	fmt.Println(log)

	// 値オブジェクトは振る舞いをもつことができる
	//   Addメソッドは変更ではなく新しい合計Moneyインスタンスを返す
	jpMoney1 := vo.NewMoney(100, "JPY")
	jpMoney2 := vo.NewMoney(200, "JPY")
	totalMoney, _ := jpMoney1.Add(jpMoney2)
	fmt.Println(fmt.Sprintf("Total: %s", totalMoney.ToString()))

	// 日本円とドルは加算できないというドメイン知識を表現する
	usMoney := vo.NewMoney(300, "USD")
	_, err := jpMoney1.Add(usMoney)
	if err != nil {
		fmt.Println(err)
	}
}
