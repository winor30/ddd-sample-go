package model

import (
	"github.com/pkg/errors"
	"github.com/winor30/ticket-service/domain"
)

type Tickets struct {
	Inventory TicketInventory
	Ticket    Ticket
}

func NewTickets(inventory TicketInventory, ticket Ticket) (*Tickets, error) {
	tickets := &Tickets{
		Inventory: inventory,
		Ticket:    ticket,
	}
	if err := tickets.validate(); err != nil {
		return nil, err
	}
	return tickets, nil
}

func (t *Tickets) Grant(count int64) (*ReducedEvent, *GrantedEvent, error) {
	if err := t.Inventory.Reduce(count); err != nil {
		return nil, nil, err
	}

	if err := t.Ticket.Granted(count); err != nil {
		// 付与に失敗した場合は、追加する
		if err := t.Inventory.Increment(count); err != nil {
			return nil, nil, err
		}
		return nil, nil, err
	}

	return newReducedEvent(t.Inventory.TicketID, count, domain.NewDomainEvent()), newGrantedEvent(t.Ticket.UserID, t.Ticket.TicketID, count, domain.NewDomainEvent()), nil
}

func (t *Tickets) validate() error {
	if t.Inventory.TicketID != t.Ticket.TicketID {
		return errors.New("inventory and ticket's ticketID is not same")
	}
	return nil
}
