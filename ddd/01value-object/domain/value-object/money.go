package valueobject

import (
	"errors"
	"fmt"
)

type Money struct {
	ammount  int
	currency string
}

func NewMoney(ammount int, currency string) *Money {
	money := &Money{
		ammount:  ammount,
		currency: currency,
	}
	return money
}

func (m Money) Equals(money *Money) bool {
	return m.currency == money.currency && m.ammount == money.ammount
}

func (m Money) Add(money *Money) (*Money, error) {
	if m.currency != money.currency {
		return nil, errors.New("通貨単位が異なります")
	}

	return NewMoney(m.ammount+money.ammount, money.currency), nil
}

func (m Money) ToString() string {
	return fmt.Sprintf("%d(%s)", m.ammount, m.currency)
}
