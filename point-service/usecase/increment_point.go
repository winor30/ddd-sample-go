package usecase

import "github.com/winor30/point-service/domain/repository"

type IncrementPointUsecase interface {
	Increment(userID string, amount int64) error
}
type incrementPointUsecase struct {
	PointAccountRepository repository.PointAccountRepository
}

func NewIncrementPointUsecase(par repository.PointAccountRepository) IncrementPointUsecase {
	return incrementPointUsecase{PointAccountRepository: par}
}

func (ipu incrementPointUsecase) Increment(userID string, amount int64) error {
	account, err := ipu.PointAccountRepository.Find(userID)
	if err != nil {
		return err
	}

	if _, err := account.Increment(amount); err != nil {
		return err
	}

	if err := ipu.PointAccountRepository.Save(account); err != nil {
		return err
	}

	return nil
}
