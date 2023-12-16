package models

import "time"

type Province struct {
	ID       uint
	Province string
}

type City struct {
	ID         uint
	City       string
	ProvinceID uint
}

type PostalCode struct {
	ID         uint
	PostalCode string
	CityID     uint
}

type Address struct {
	ID           uint
	ProvinceID   uint
	CityID       uint
	PostalCodeID uint
	Address      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
