package responses

import "time"

type UserResponse struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Birthdate   string    `json:"birthdate"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserLogin struct {
	Token string `json:"token"`
}

type AddressDetail struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Province   string `json:"province"`
}

type UserDetail struct {
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	PhoneNumber string          `json:"phone_number"`
	Birthdate   string          `json:"birthdate"`
	Addresses   []AddressDetail `json:"addresses"`
}
