package model

import (
	"github.com/winor30/ticket-service/domain"
)

type GrantedEvent struct {
	UserID   string
	TicketID string
	Count    int64
	*domain.Event
}

func newGrantedEvent(userID, ticketID string, count int64, event *domain.Event) *GrantedEvent {
	return &GrantedEvent{
		Event:    event,
		UserID:   userID,
		TicketID: ticketID,
		Count:    count,
	}
}
