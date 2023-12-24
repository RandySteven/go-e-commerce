package requests

import "github.com/shopspring/decimal"

type ProductRequest struct {
	Name        string          `json:"name" validate:"required,min=8,max=54"`
	Price       decimal.Decimal `json:"price" validate:"required,numeric"`
	Stock       uint            `json:"stock" validate:"numeric"`
	Description string          `json:"description" validate:"max=192"`
	CategoryID  uint            `json:"category_id" validate:"required"`
}
