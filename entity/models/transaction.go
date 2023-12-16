package models

import "time"

type Transaction struct {
	ID              uint
	ShopID          uint
	UserID          uint
	TransactionCode string
	TransactionDate time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
