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
