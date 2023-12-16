package requests

type UserRegisterRequest struct {
	FirstName   string `json:"first_name" validate:"required,min=3,max=12"`
	LastName    string `json:"last_name" validate:"required,min=3,max=12"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,alphanum,min=8,max=16"`
	PhoneNumber string `json:"phone_number" validate:"required,numeric,startWith=+62,min=10,max=12"`
	Birthday    string `json:"birthdate" validate:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum,min=8,max=16"`
}
