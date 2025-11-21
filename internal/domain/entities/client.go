package entities

import "time"

type Client struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	DocumentType string    `json:"document_type"`
	DocumentID   string    `json:"document_id"`
	Status       string    `json:"status"` //active, blocked, pending
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c Client) Validate() error {
	return nil
}

func (c Client) IsActive() bool {
	return c.Status == "active"
}

func (c Client) GetFullName() string {
	return c.FirstName + " " + c.LastName
}
