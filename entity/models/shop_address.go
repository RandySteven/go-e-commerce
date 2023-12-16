package models

import "time"

type ShopAddress struct {
	ID        uint
	AddressID uint
	ShopID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletdAt  *time.Time
	Address   Address
	Shop      Shop
}
