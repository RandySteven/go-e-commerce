package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          uint
	Name        string
	Price       decimal.Decimal
	Stock       uint
	CategoryID  uint
	ShopID      uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Category    *Category
	Shop        *Shop
}
