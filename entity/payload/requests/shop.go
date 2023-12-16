package requests

type ShopRegisterRequest struct {
	Name        string `json:"name" validate:"required,min=8,max=64"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,alphanum,min=8,max=16"`
	PhoneNumber string `json:"phone_number" validate:"required,numeric,startWith=+62,min=10,max=12"`
}

type ShopLoginRequest struct {
	Name     string `json:"name" validate:"min=8,max=64"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,alphanum,min=8,max=16"`
}
