package model

import (
	"github.com/winor30/point-service/domain"
)

type IncrementEvent struct {
	Amount  int64
	Balance int64
	*domain.Event
}

func newIncrementEvent(amount, balance int64, event *domain.Event) *IncrementEvent {
	return &IncrementEvent{
		Event:   event,
		Amount:  amount,
		Balance: balance,
	}
}
