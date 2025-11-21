package entities

import "time"

type Wallet struct {
	ID        string    `json:"id"`
	ClientID  string    `json:"client_id"`
	Currency  string    `json:"currency"` // RUB, USD, EUR
	Balance   float64   `json:"balance"`
	Status    string    `json:"status"` //active, frozen, closed
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CanWithdraw(amount float64) bool {
	return false
}

func (w Wallet) CanDeposit(amount float64) bool {
	return false
}

func (w Wallet) GetBalance() float64 {
	return w.Balance
}

func (w Wallet) IsActive() bool {
	return w.Status == "active"
}

func (w Wallet) Freeze() error {
	return nil
}

func (w Wallet) Unfreeze() error {
	return nil
}

func (w Wallet) Close() error {
	return nil
}
