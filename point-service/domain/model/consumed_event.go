package model

import (
	"github.com/winor30/point-service/domain"
)

type ConsumedEvent struct {
	Amount  int64
	Balance int64
	*domain.Event
}

func newConsumedEvent(amount, balance int64, event *domain.Event) *ConsumedEvent {
	return &ConsumedEvent{
		Event:   event,
		Amount:  amount,
		Balance: balance,
	}
}
