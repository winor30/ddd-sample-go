package model

import "github.com/pkg/errors"

type Ticket struct {
	UserID   string
	TicketID string
	Count    int64
}

func (t *Ticket) Granted(count int64) error {
	if count <= 0 {
		return errors.New("granted count is less than zero")
	}
	t.Count = count
	return nil
}
