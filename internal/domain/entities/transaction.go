package entities

import "time"

type Transaction struct {
	ID           string    `json:"id"`
	Type         string    `json:"type"` // deposit, withdrawal, transfer
	FromWalletID string    `json:"from_wallet_id,omitempty"`
	ToWalletID   string    `json:"to_wallet_id,omitempty"`
	Amount       float64   `json:"amount"`
	Currency     string    `json:"currensy"`
	Status       string    `json:"status"` // pending, completed, failed, cancelled
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	CompletedAt  time.Time `json:"completed_at,omitempty"`
}

func (t Transaction) Validate() error {
	return nil
}

func (t Transaction) CanProcess() bool {
	return false
}

func (t Transaction) MarkAsCompleted() error {
	return nil
}

func (t Transaction) MarkAsFailed() error {
	return nil
}

func (t Transaction) GetTransactionDetails() string {
	return ""
}
