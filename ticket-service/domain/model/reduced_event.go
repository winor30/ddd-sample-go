package model

import (
	"github.com/winor30/ticket-service/domain"
)

type ReducedEvent struct {
	TicketID string
	Count    int64
	*domain.Event
}

func newReducedEvent(ticketID string, count int64, event *domain.Event) *ReducedEvent {
	return &ReducedEvent{
		Event:    event,
		TicketID: ticketID,
		Count:    count,
	}
}
