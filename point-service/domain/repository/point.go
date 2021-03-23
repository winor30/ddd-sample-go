package repository

import (
	"github.com/winor30/point-service/domain/model"
)

type PointAccountRepository interface {
	Find(userID string) (*model.PointAccount, error)
	Save(account *model.PointAccount) error
}
