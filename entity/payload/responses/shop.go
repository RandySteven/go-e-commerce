package responses

type ShopResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type ShopLoginResponse struct {
	Token string `json:"token"`
}
