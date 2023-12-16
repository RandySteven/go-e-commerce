package requests

type UserRegisterRequest struct {
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Birthday    string
}

type UserLoginRequest struct {
	Email    string
	Password string
}
