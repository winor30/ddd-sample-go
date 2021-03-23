package usecase

import "github.com/winor30/point-service/domain/repository"

type ConsumePointUsecase interface {
	Consume(userID string, amount int64) error
}
type consumePointUsecase struct {
	PointAccountRepository repository.PointAccountRepository
}

func NewConsumePointUsecase(par repository.PointAccountRepository) ConsumePointUsecase {
	return consumePointUsecase{PointAccountRepository: par}
}

func (ipu consumePointUsecase) Consume(userID string, amount int64) error {
	account, err := ipu.PointAccountRepository.Find(userID)
	if err != nil {
		return err
	}

	if _, err := account.Consume(amount); err != nil {
		return err
	}

	if err := ipu.PointAccountRepository.Save(account); err != nil {
		return err
	}

	return nil
}
