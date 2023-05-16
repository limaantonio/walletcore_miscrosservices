package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) *Account {
	if client == nil {
		return nil
	}
	account := &Account{
		Id:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := account.Validate()
	if err != nil {
		return nil
	}

	return account
}

func (a *Account) Credit(amount float64) {
	a.Balance += amount
	a.UpdatedAt = time.Now()
}

func (a *Account) Debit(amount float64) {
	a.Balance -= amount
	a.UpdatedAt = time.Now()
}

func (a *Account) Validate() error {
	if a.Client == nil {
		return errors.New("client is required")
	}
	return nil
}
