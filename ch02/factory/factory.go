package factory

import (
	"errors"
	"fmt"
)

const (
	Cash      = 1
	DebitCard = 2
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New("Not implemented yet")
	}
}

type CashPM struct{}

type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%.2f paid using cash", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%.2f paid using debit card", amount)
}
