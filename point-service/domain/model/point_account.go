package model

import (
	"time"

	"github.com/pkg/errors"
	"github.com/winor30/point-service/domain"
)

type PointAccount struct {
	UserID    string
	Balance   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(userID string, balance int64, createdAt, updatedAt time.Time) *PointAccount {
	return &PointAccount{
		UserID:    userID,
		Balance:   balance,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (a *PointAccount) Increment(amount int64) (*IncrementEvent, error) {
	if amount <= 0 {
		return nil, errors.New("amount is less than zero")
	}

	a.Balance = a.Balance + amount
	a.updateTime()
	return newIncrementEvent(amount, a.Balance, domain.NewDomainEvent()), nil
}

func (a *PointAccount) Consume(amount int64) (*ConsumedEvent, error) {
	if amount <= 0 {
		return nil, errors.New("amount is less than zero")
	}

	newBalance := a.Balance - amount
	if newBalance <= 0 {
		return nil, errors.Errorf("insufficient balance. balance: %d, amount: %d", a.Balance, amount)
	}

	a.Balance = newBalance
	a.updateTime()
	return newConsumedEvent(amount, a.Balance, domain.NewDomainEvent()), nil
}

func (a *PointAccount) updateTime() {
	a.UpdatedAt = time.Now()
}
