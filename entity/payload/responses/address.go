package responses

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

type ShopDetail struct {
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	PhoneNumber string        `json:"phone_number"`
	Address     AddressDetail `json:"address"`
}
