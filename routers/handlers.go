package routers

import (
	"github.com/RandySteven/go-e-commerce.git/db/postgres"
	"github.com/RandySteven/go-e-commerce.git/handlers"
	"github.com/RandySteven/go-e-commerce.git/interfaces"
	"github.com/RandySteven/go-e-commerce.git/usecases"
)

type Handlers struct {
	interfaces.UserHandler
	interfaces.ShopHandler
}

func NewHandlers(repo *postgres.Repositories) *Handlers {
	userUsecase := usecases.NewUserUsecase(repo.UserRepository)
	shopUsecase := usecases.NewShopUsecase(repo.ShopRepository)

	return &Handlers{
		UserHandler: handlers.NewUserHandler(userUsecase),
		ShopHandler: handlers.NewShopHandler(shopUsecase),
	}
}
