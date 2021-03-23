package usecase

import "github.com/winor30/ticket-service/domain/repository"

type GrantTicketUsecase interface {
	Grant(userID, ticketID string, count int64) error
}
type grantTicketUsecase struct {
	TicketsRepository repository.TicketsRepository
}

func NewConsumePointUsecase(tr repository.TicketsRepository) GrantTicketUsecase {
	return &grantTicketUsecase{TicketsRepository: tr}
}

func (gtu *grantTicketUsecase) Grant(userID, ticketID string, count int64) error {
	tickets, err := gtu.TicketsRepository.Find(userID, ticketID)
	if err != nil {
		return err
	}

	if _, _, err := tickets.Grant(count); err != nil {
		return err
	}

	if err := gtu.TicketsRepository.Save(tickets); err != nil {
		return err
	}

	return nil
}
