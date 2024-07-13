package model

import "time"

type Transaction struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	InvoiceNumber string    `json:"invoice_number"`
	Status        string    `json:"status"`
	Amount        int       `json:"amount"`
	Items         string    `json:"items"`
	InvoiceDate   string    `json:"invoice_date"`
	UserID        uint      `json:"user_id"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
