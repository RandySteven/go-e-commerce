package models

import "time"

type Shop struct {
	ID          uint
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
