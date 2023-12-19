package handlers

import "github.com/RandySteven/go-e-commerce.git/interfaces"

type ShopHandler struct {
	usecase interfaces.ShopUsecase
}

func NewShopHandler(usecase interfaces.ShopUsecase) *ShopHandler {
	return &ShopHandler{
		usecase: usecase,
	}
}
