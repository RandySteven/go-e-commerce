package models

import "time"

type Cart struct {
	ID        uint
	ShopID    uint
	UserID    uint
	ProductID uint
	Quantity  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
