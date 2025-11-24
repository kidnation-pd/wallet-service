package repositories

import (
	"time"
	"wallet/internal/domain/entities"
)

type ClientRepository interface {
	Create(client *entities.Client) error
	GetByID(id string) (*entities.Client, error)
	GetByEmail(email string) (*entities.Client, error)
	GetByDocument(documentID string) (*entities.Client, error)
	Update(client *entities.Client) error
	Delete(id string) error
	List(limit, offset int) ([]*entities.Client, error)
}
type WalletRepository interface {
	Create(wallet *entities.Wallet) error
	GetByID(id string) (*entities.Wallet, error)
	GetClientID(clientID string) (*[]entities.Wallet, error)
	Update(wallet *entities.Wallet) error
	UpdateBalance(walletID string, newBalance float64) error
	Delete(id string) error
}

type TransactionRepository interface {
	Create(transaction *entities.Transaction) error
	GetByID(id string) (*entities.Transaction, error)
	GetByWalletID(walletID string, limit, offset int) ([]*entities.Transaction, error)
	Update(transaction *entities.Transaction) error
	GetTransactionsByPeriod(walletID string, from, to time.Time) ([]*entities.Transaction, error)
}
