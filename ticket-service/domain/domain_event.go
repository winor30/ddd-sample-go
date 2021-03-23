package domain

import (
	"time"
)

type Event struct {
	ID        string
	CreatedAt time.Time
}

func NewDomainEvent() *Event {
	return &Event{
		ID:        randString(32),
		CreatedAt: time.Time{},
	}
}
