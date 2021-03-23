package model

import "github.com/pkg/errors"

type TicketInventory struct {
	TicketID string
	Stock    int64
}

func (i *TicketInventory) Reduce(value int64) error {
	if value <= 0 {
		return errors.New("inventory reduce count is less than zero")
	}
	newStock := i.Stock - value
	if newStock <= 0 {
		return errors.New("inventory reduce count is greater than stock")
	}

	i.Stock = newStock
	return nil
}

func (i *TicketInventory) Increment(value int64) error {
	if value <= 0 {
		return errors.New("inventory reduce count is less than zero")
	}
	newStock := i.Stock + value

	i.Stock = newStock
	return nil
}
