package models

import "time"

type Category struct {
	ID        uint
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
