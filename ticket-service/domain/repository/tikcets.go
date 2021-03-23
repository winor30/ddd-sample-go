package repository

import (
	"github.com/winor30/ticket-service/domain/model"
)

type TicketsRepository interface {
	Find(userID, ticketID string) (*model.Tickets, error)
	Save(tickets *model.Tickets) error
}
