package models

import "time"

type UserAddress struct {
	ID        uint
	AddressID uint
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
